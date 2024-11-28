package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	fmt.Println("spawn goroutine export metrics")
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()
	fmt.Println("spawn goroutine infinite loop")
	go func(){
		for{}
	}()
	fmt.Println("run infinite loop")
	for {}
}