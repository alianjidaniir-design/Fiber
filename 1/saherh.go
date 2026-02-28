package main

import (
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"path"
	"time"
)

func myprint(start, finish int) {
	for i := 0; i <= finish; i++ {
		fmt.Println(i, "Amir Al_Momenin")
	}
	fmt.Println("")
	time.Sleep(2 * time.Second)
}

func main() {

	for i := 0; i < 7; i++ {
		go myprint(i, 9)
	}

	sys, err := syslog.New(syslog.LOG_SYSLOG, "saherh.go")
	fmt.Println("Hello")
	if err != nil {
		log.Println("Error creating syslog")
		return
	}

	time.Sleep(2 * time.Second)

	LOGFILE := path.Join(os.TempDir(), "saherh.log")
	fmt.Println(LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error opening log file")
		return
	}
	defer f.Close()

	LsstdFlags := log.Lshortfile | log.Ldate
	iLog := log.New(f, "LNum ", LsstdFlags)
	iLog.Println("Mastering Go, 4th edition!")

	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another log entry!")

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

}
