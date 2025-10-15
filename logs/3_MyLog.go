package main

import (
	"log"
	"os"
)

func main() {
	flog, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) 
	// os.O_WRONLY — открыть файл только для записи (write-only).
	// 0644 — это права доступа в Unix-подобной системе (восьмеричное число).
	// Владелец: может читать и писать (rw-)
	// Группа: только читать (r--)
	// Остальные: только читать (r--)
	// То есть в символах это -rw-r--r--.

	// flog — это *os.File, то есть открытый файловый дескриптор.
	if err != nil {
		log.Fatal(err)
	}

	defer flog.Close()

	mylog := log.New(flog, "serv ", log.LstdFlags|log.Lshortfile)
	// flog — это io.Writer, куда будет идти лог.
	// "serv " — это префикс для каждой строки лога.
	// log.LstdFlags | log.Lshortfile — комбинация флагов форматирования.
	// serv 2025/10/15 14:21:50 3_MyLog.go:18: Start server 

	mylog.Println("Start server")
	mylog.Println("Fineshed server")

}