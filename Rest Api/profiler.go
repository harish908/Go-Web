package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "net/http/pprof"
)

// Some function that does work
func hardWork(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Start: %v\n", time.Now())

	// Memory
	a := []string{}
	for i := 0; i < 500000; i++ {
		a = append(a, "aaaa")
	}

	// Blocking
	time.Sleep(2 * time.Second)
	fmt.Printf("End: %v\n", time.Now())
}

// run program and run "go tool pprof http://localhost:8000/debug/pprof/heap"
func main() {
	var wg sync.WaitGroup

	// Server for pprof
	go func() {
		fmt.Println(http.ListenAndServe("localhost:8000", nil))
	}()
	wg.Add(1) // pprof - so we won't exit prematurely
	wg.Add(1) // for the hardWork
	go hardWork(&wg)
	wg.Wait()
}