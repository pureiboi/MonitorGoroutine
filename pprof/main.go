package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	fmt.Println("spawn goroutine profiling server")
	go func() {
		// localhost:6060/debug/pprof/goroutine?debug=2
		http.ListenAndServe("localhost:6060", nil)
	}()

	fmt.Println("spawn goroutine infinite loop")
	go func() {
		for {}
	}()

	fmt.Println("run infinite loop")
	for {}
}
