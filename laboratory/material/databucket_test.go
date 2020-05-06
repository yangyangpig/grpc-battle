package material

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestMyDataBucket(t *testing.T) {
	dataBucket := NewMyDataBucket()

	//c := make(chan struct{})
	//for i:=0; i < 10; i++ {
	//	go dataBucket.Read(i)
	//}
	go dataBucket.Read(1)
	go dataBucket.Read(2)

	for i:=0; i < 10; i++ {
		go func(i int) {
			d := fmt.Sprintf("data-%d", i)
			dataBucket.Put([]byte(d))
		}(i)
		//time.Sleep(100 * time.Millisecond)

	}

	select {
	case <-time.After(time.Second * 2):

	}


}

func TestMyDataBucket_Put(t *testing.T) {
	a := "1587638318768160596"
	msInt, _ := strconv.ParseInt(a, 10, 64)

	tm := time.Unix(0, msInt*int64(time.Nanosecond))

	fmt.Println(tm.Format("2006-02-01 15:04:05.000"))

}