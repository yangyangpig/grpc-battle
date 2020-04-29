package laboratory

// 等待put协程通知，说池子有数据了，才开始读，不然会一直在那等着,这个动作可以实现无论上面有多少个goroutine在读，但是都是顺序读出来的
// 因为有数据才读出来，达到不会因为多个goroutine读取造成独处数据乱序问题

// 实现了，多个goroutine生产数据，多个goutine消费数据，但是，读出来的数据都还是有序的

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

type MyDataBucket struct {
	br *bytes.Buffer
	gmutex *sync.RWMutex
	rcond *sync.Cond
}

func NewMyDataBucket() *MyDataBucket {
	buf := make([]byte, 0)
	db := &MyDataBucket{
		br:     bytes.NewBuffer(buf),
		gmutex: new(sync.RWMutex),
	}

	db.rcond = sync.NewCond(db.gmutex.RLocker())
	return db
}

func (b *MyDataBucket) Read(r int) {
	b.gmutex.RLock()
	defer b.gmutex.RUnlock()

	var data []byte
	var d byte
	var err error

	for {
		if d, err = b.br.ReadByte(); err != nil {
			if err == io.EOF {
				if string(data) != "" {
					fmt.Printf("reader-%d: %s\n", r, data)
				}
				// 等待put协程通知，说池子有数据了，才开始读，不然会一直在那等着,这个动作可以实现无论上面有多少个goroutine在读，但是都是顺序读出来的
				// 因为有数据才读出来，达到不会因为多个goroutine读取造成独处数据乱序问题
				b.rcond.Wait()
				// 清空data
				data = data[:0]
				continue
			}
		}
		data = append(data, d)
		//fmt.Printf("reader-%d the bucket data : %+v\n", r, data)

	}
}

func (b *MyDataBucket) Put(d []byte) (int, error) {
	b.gmutex.Lock()
	defer  b.gmutex.Unlock()
	n, err := b.br.Write(d)
	b.rcond.Signal()
	return n, err
}


