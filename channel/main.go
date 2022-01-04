package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch = make(chan int, 10)
var ch_done = make(chan int)
var NUM = 15

func write() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < NUM; i++ {
		num := rand.Intn(1000)
		ch <- num
		time.Sleep(1 * time.Second)
	}
}

func read() {
	i := 1
	for {
		num := <-ch
		fmt.Println(num)
		i += 1
		if i > NUM {
			ch_done <- 0
		}
	}
}

func main() {
	go write()
	go read()
	done := <-ch_done
	fmt.Println("Done!", done)
}
