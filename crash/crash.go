package crash

import (
	"log"
	"runtime/debug"
)

func nested() {
	var mySlice []int
	j := mySlice[0] // crash
	log.Printf("Hello: %d\n", j)
}

func Crash() {
	debug.SetPanicOnFault(true)
	debug.SetTraceback("all")

	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered in goroutine")
			stackTrace := string(debug.Stack())
			log.Println(stackTrace)
		}
	}()

	nested()
}
