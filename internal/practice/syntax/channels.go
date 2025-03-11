package syntax

import "fmt"

func Channels() {
	pings := make(chan string)
	pongs := make(chan string)

	go pinger(pings)
	go ponger(pings, pongs)
	go printer(pongs)
}

func pinger(pings chan<- string) {
	for i := 0; i < 5; i++ {
		pings <- "ping"
	}
}

func ponger(pings <-chan string, pongs chan<- string) {
	for {
		select {
		case msg := <-pings:
			pongs <- msg
		default:
			pongs <- "pong"
		}
	}
}

func printer(pongs <-chan string) {
	for {
		pong := <-pongs
		fmt.Println(pong)
	}
}
