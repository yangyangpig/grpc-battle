package main

import (
	"fmt"
	"grpc-battle/laboratory/material"
	"time"
)

func main()  {
	dataBucket := material.NewMyDataBucket()

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
