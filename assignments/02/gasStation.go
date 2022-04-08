package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func pump(fuelType string, wg *sync.WaitGroup, minMilSecsWait, maxMilSecsWait int, pumpCap chan int) {
	defer wg.Done()
	cashDesk := make(chan int, 2)
	carQueue := 3
	var wg2 sync.WaitGroup
	for i := 0; i < carQueue; i++ {
		pumpCap <- 1
		wg2.Add(1)
		waitTime := rand.Intn(maxMilSecsWait-minMilSecsWait) + minMilSecsWait
		go func() {
			cashDesk <- 1
			carServe(fuelType, waitTime, &wg2)
			<-pumpCap
			<-cashDesk
		}()
	}
	wg2.Wait()
}

func carServe(fuelT string, waitTime int, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	fmt.Println("Car with fuel", fuelT, "fueling..")
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	fmt.Println("Car with fuel", fuelT, "paying..")
	time.Sleep(time.Duration(rand.Intn(2-1)+1) * time.Millisecond)
}

func main() {
	pumpCap := make(chan int, 4)
	elCap := make(chan int, 8)
	lpgCap := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go pump("gas", &wg, 1000, 5000, pumpCap)
	wg.Add(1)
	go pump("diesel", &wg, 1000, 5000, pumpCap)
	wg.Add(1)
	go pump("lpg", &wg, 1000, 5000, lpgCap)
	wg.Add(1)
	go pump("electric", &wg, 3000, 10000, elCap)
	wg.Wait()
}
