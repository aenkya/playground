package syntax

import (
	"fmt"
	"time"
)

func goroutines() {
	quit := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("Hello")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	quit <- true
}
