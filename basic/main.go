package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
}
