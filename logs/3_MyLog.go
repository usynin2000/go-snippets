package main

import (
	"log"
	"os"
)

func main() {
	flog, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// os.O_WRONLY — open file only for writing (write-only).
	// 0644 — this is access rights in Unix-like system (octal number).
	// Owner: can read and write (rw-)
	// Group: can only read (r--)
	// Others: can only read (r--)
	// So in symbols it is -rw-r--r--.

	// flog — this is *os.File, so opened file descriptor.
	if err != nil {
		log.Fatal(err)
	}

	defer flog.Close()

	mylog := log.New(flog, "serv ", log.LstdFlags|log.Lshortfile)
	// flog — это io.Writer, куда будет идти лог.
	// "serv " — this is prefix for each log line.
	// log.LstdFlags | log.Lshortfile — combination of formatting flags.
	// serv 2025/10/15 14:21:50 3_MyLog.go:18: Start server

	mylog.Println("Start server")
	mylog.Println("Fineshed server")

}
