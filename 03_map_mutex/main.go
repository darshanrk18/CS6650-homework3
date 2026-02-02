package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func main() {
	s := SafeMap{m: make(map[int]int)}
	var wg sync.WaitGroup

	const goroutines = 50
	const perG = 1000

	start := time.Now()

	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func(g int) {
			defer wg.Done()
			for i := 0; i < perG; i++ {
				key := g*perG + i
				s.mu.Lock()
				s.m[key] = i
				s.mu.Unlock()
			}
		}(g)
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println("len(m):", len(s.m))
	fmt.Println("time  :", elapsed)
}