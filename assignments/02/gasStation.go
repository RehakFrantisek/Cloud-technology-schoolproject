package main

import (
	"fmt"
	"sync"
)

func pump(spots int, fuel_type string, group *sync.WaitGroup) {
	var wg2 sync.WaitGroup
	for i := 0; i < spots; i++ {
		wg2.Add(1)
		go func(n int) {
			fmt.Println("test")
		}(i)
		wg2.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go pump(4, "gas", &wg)
	wg.Add(1)
	go pump(4, "diesel", &wg)
	wg.Add(1)
	go pump(1, "lpg", &wg)
	wg.Add(1)
	go pump(8, "electric", &wg)
}
