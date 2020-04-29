package laboratory

import (
	"fmt"
	"testing"
	"time"
)

func TestMyDataBucket(t *testing.T) {
	dataBucket := NewMyDataBucket()

	//c := make(chan struct{})
	for i:=0; i < 10; i++ {
		go dataBucket.Read(i)
	}

	for i:=0; i < 10; i++ {
		go func(i int) {
			d := fmt.Sprintf("data-%d", i)
			dataBucket.Put([]byte(d))
		}(i)
		time.Sleep(100 * time.Millisecond)

	}


}