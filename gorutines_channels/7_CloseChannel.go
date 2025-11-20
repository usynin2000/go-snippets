package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Exmaple 1: Closing channel and reading")
	ch := make(chan int, 3)

	// Sending data
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // closing channel

	// Reading all data
	
	fmt.Println("Reading form closed channel:")
	for i := 0; i < 4; i++ {
		value, ok := <-ch
		if !ok {
			fmt.Println(" Channel is closed or empty")
			break
		}
		fmt.Printf("got %d\n", value)
	} // If not closing will get an error all goroutines are asleep - deadlock!

	fmt.Println("\n=== Example 2: range automatically handling closing ===")
	ch2 := make(chan string)

	go func() {
		ch2 <- "first"
		ch2 <- "second"
		ch2 <- "third"
		close(ch2)
	}()

	for msg := range ch2 {
		fmt.Printf("Got %s,\n", msg)
	}

	fmt.Println("Cycle stopped automatically after channel closing")


	fmt.Println("\n === Example 3: Gorutines are notified about channels")
	ch3 := make(chan int)
	done := make (chan bool)

	go func() {
		for {
			value, ok := <-ch3
			if !ok {
				fmt.Println("Gorutine got a signal about closing channel")
				done <- true
				return
			}
			fmt.Printf("Gorutine handled: %d\n", value)
		}
	}()

	// sending data
	ch3 <- 10
	ch3 <- 20
	time.Sleep(500 * time.Millisecond)

	// Closing channel = signal to gorutine there is no data
	close(ch3)
	<-done // Waiting confirmation from gorutine

	fmt.Println("\n === NB! Closing channel == signal NO DATA ANYMORE ===")


}


// Что показывает:
// Как закрывать каналы
// Проверка закрытия через ok
// range автоматически завершается при закрытии
// Закрытие как сигнал горутинам
