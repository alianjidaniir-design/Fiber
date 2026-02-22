package main

import (
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
)

func main() {
	sys, err := syslog.New(syslog.LOG_SYSLOG, "saherh.go")
	fmt.Println("Hello")
	if err != nil {
		log.Println("Error creating syslog")
		return
	}

	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	fiel, err := os.OpenFile("saherh.log", flag, 0664)
	if err != nil {
		log.Println("Error opening file")
		return
	}
	w := io.MultiWriter(fiel, os.Stdout, os.Stderr)
	logger := log.New(w, "", log.LstdFlags|log.Lshortfile)
	logger.Println("Hello")
	logger.Println("BOOK %d", os.Getpid())

	log.Println(sys)
	log.Println("Everything is good")

	//saeed

}
