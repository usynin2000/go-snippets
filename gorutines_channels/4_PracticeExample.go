package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simulating loading files
func downloadFile(url string, ch chan <- string) {
	// simulating time downloading (1-3 sec)
	duration := time.Duration(rand.Intn(2000)+ 1000) * time.Millisecond
	time.Sleep(duration)

	ch <- fmt.Sprintf("File %s was uploaded in %v milliseconds", url, duration)
}

func main() {
	fmt.Println("=== Loading without gorutines (sequentially) ===")
	start := time.Now()

	urls := []string{"file1.pdf", "file2.pdf", "file3.pdf", "fil4.pdf"}

	for _, url := range urls {
		fmt.Printf("Loading %s...\n", url)
		time.Sleep(1500 * time.Millisecond)
		fmt.Printf("+ %s loaded\n", url)
	}

	elapsed := time.Since(start)
	fmt.Printf("Time passed %v\n\n", elapsed)

	fmt.Println("=== Loading with Gorutines (in parallel) ===")
	start = time.Now()
	ch := make(chan string, len(urls))

	// Start to load in parallel
	for _, url := range urls {
		go downloadFile(url, ch)
	}

	// Returning result
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}

	elapsed = time.Since(start)
	fmt.Printf("Time passed %v\n\n", elapsed)
	fmt.Println("Look at the difference! Downloading with gorutines much faster!")

}


// Что показывает:
// Зачем нужны горутины: ускорение параллельных операций
// Практический случай: загрузка нескольких файлов
// Сравнение последовательного и параллельного выполнения