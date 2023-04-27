package main

import (
	"log"
	"os"
	"recoverTest/crash"
	"runtime/debug"
	"time"
)

func main() {
	debug.SetPanicOnFault(true)
	debug.SetTraceback("all")

	// setting up logs
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)

	// capturing panic details in main thread won't work
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered in main thread")
			log.Println(string(debug.Stack()))
		}
	}()

	go crash.Crash() // crashes

	time.Sleep(1 * time.Second)

	_ = f.Close()
}
