package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func main() {
	var c1, c2 = generator(), generator()
	for {
		select {
		case n := <-c1:
			fmt.Printf("received from c1...%d \n", n)
		case n := <-c2:
			fmt.Printf("received from c1...%d \n", n)
			//default:
			//	fmt.Println("no value received...")
		}
	}

}
