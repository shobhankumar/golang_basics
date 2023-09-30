package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex

	var sslice = make([]int, 10)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("calling from go routine 1")
		mut.Lock()
		sslice = append(sslice, 100)
		mut.Unlock()
		wg.Done()
	}(&wg, &mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("calling from go routine 2")
		mut.Lock()
		sslice = append(sslice, 200)
		mut.Unlock()
		wg.Done()
	}(&wg, &mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("calling from go routine 3")
		mut.Lock()
		sslice = append(sslice, 300)
		mut.Unlock()
		wg.Done()
	}(&wg, &mut)

	wg.Wait()

	fmt.Println(sslice)
}
