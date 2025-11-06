package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("=== Example 1: Unbuffered Channel ===")
	unbuffered := make(chan int)

	go func() {
		fmt.Println("Try to send data to unbuffered channel...")
		unbuffered <- 1 // it is bloacked unless somebody read it
		fmt.Println("Sent! (it will run after somebody read from unbuffered)")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Reading from unbuffered channel..")
	value := <-unbuffered
	fmt.Printf("Got value form unbuffered channel: %d\n\n", value)


	fmt.Println("=== Exmaple 2: Buffered channel ===")
	buffered := make(chan int, 3) // buffer for 3 elements capacity
	
	go func() {
		// We can sen values without block
		fmt.Println("Sending 3 values to buffer...")
		buffered <- 1
		buffered <- 2
		buffered <- 3
		fmt.Println("All three values was sent WITHOUT blocking!\n")
	}()
	
	// reading
	time.Sleep(500 * time.Millisecond) // give it some time to send data
	fmt.Println("Reading from buffer:")
	fmt.Printf(" Got %d\n", <-buffered)
	fmt.Printf(" Got %d\n", <-buffered)
	fmt.Printf(" Got %d\n", <-buffered)
	fmt.Println()


	fmt.Println("=== Example 3: Buffer Overflow ===")
	smallBuffer := make(chan int, 3)

	go func() {
		for i := 1; i <= 4; i++ {
			fmt.Printf("Sending %d...\n", i)
			smallBuffer <- i // It will be blocked in 4t element
			fmt.Printf("%d sent\n", i)
		}
		close(smallBuffer)
	}()

	time.Sleep(2 * time.Second) // Give it some time to send first 2 digits
	fmt.Println("Started to read...")
	for val := range smallBuffer {
		fmt.Printf("Read: %d\n", val)
		time.Sleep(500 * time.Millisecond)
	}

}


// Пример 1: Небуферизованный канал (Unbuffered Channel)

// - Горутина пытается отправить значение в канал
// - Операция unbuffered <- 1 блокируется до тех пор, пока кто-то не прочитает
// - Горутина ждет, пока main не прочитает значение
// - Только после чтения выполняется следующая строка "Sent!"


// - Ключевая идея: небуферизованный канал = синхронизация. Отправитель и получатель должны встретиться одновременно.


// Пример 2: Буферизованный канал (Buffered Channel)

// - Создается канал с буфером на 3 элемента: make(chan int, 3)
// - Горутина может отправить 3 значения подряд без блокировки
// - Значения сохраняются в буфере
// - Main читает их позже

// - Ключевая идея: буфер = очередь. Можно отправить несколько значений без ожидания читателя.


// Пример 3: Переполнение буфера (Buffer Overflow)

// - Буфер на 3 элемента, отправляется 4 значения
// - Первые 3 уходят в буфер без блокировки
// - 4-е значение блокирует горутину, пока main не прочитает хотя бы одно значение
// - После чтения одного значения горутина может отправить 4-е

// Ключевая идея: при переполнении буфера отправка блокируется, пока не освободится место.


// Итог
// Небуферизованные каналы: синхронизация, гарантия обработки
// Буферизованные каналы: производительность, асинхронность, очередь
// Выбор зависит от задачи: нужна ли гарантия обработки или важнее скорость отправки.