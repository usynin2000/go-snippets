package main

import "fmt"


func main() {
	fmt.Println("Reading from channel\n")

	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i // add to channel value
	}
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}

	ch1 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch1 <- i // add to channel value
	}

	ch2 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch2 <- i // add to channel value
	}
	close(ch2)

	fmt.Println("\nReading from channels\n")

	// It is very useful when we need to listen all channels
	for {
		select {
		case v, ok := <- ch2:
			if !ok {
				return // channel is closed
			}
			fmt.Println("from ch2:", v)
		case v, ok := <- ch1:
			if !ok {
				return
			}
			fmt.Println("from ch1:", v)
		}
	}
}
