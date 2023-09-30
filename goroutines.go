// This file show cases how to use following
// go routines
// anonymous function
// Waitgroups
// Mutexs
// mutexes are need as below example has a slice thats modified by various go routines (i.e threads)
// since slice is being appened by multiple goroutines, need to safegaurd slice using a mutex
// when we have goroutines, waitgroups also needed as main thread can exit before even
// goroutines start executing

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
