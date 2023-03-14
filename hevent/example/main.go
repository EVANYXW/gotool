package main

import (
	"demo/tools/hevent"
	"fmt"
	"time"
)

func main() {
	var ch = make(chan hevent.HEvent)
	var ch2 = make(chan hevent.HEvent)
	hevent.HEventSrv().Sub("hw", ch)
	hevent.HEventSrv().Sub("hw", ch2)
	//go GetHEventSrv().Push("topic1", "Hi topic 1")
	go func() {
		for {
			time.Sleep(1 * time.Second)
			hevent.HEventSrv().Push("hw", "Hi topic 1")
		}
	}()
	for {
		select {
		case c := <-ch:
			fmt.Println("ch", c)
		case c := <-ch2:
			fmt.Println("ch2", c)
		}
	}
}
