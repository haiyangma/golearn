package persistent

import "fmt"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			fmt.Printf("Got item : %v \n", item)
		}
	}()
	return out
}
