package main

import (
	"fmt"
	"sync"
)

func createCashier(cashierID int, wg *sync.WaitGroup, mut sync.Mutex) func(int) {
	ordersProcessed := 0
	return func(orderNum int) {
		mut.Lock()
		if ordersProcessed < 10 {

			fmt.Println(cashierID, "->", ordersProcessed)
			ordersProcessed++
		} else {
			fmt.Println("Cashier ", cashierID, "I am tired! I want to take rest!", orderNum)
		}
		wg.Done()
		mut.Unlock()
	}
}

func main() {
	cashierIndex := 0
	var wg sync.WaitGroup
	var w sync.Mutex
	cashiers := []func(int){}
	for i := 1; i <= 3; i++ {
		cashiers = append(cashiers, createCashier(i, &wg, w))
	}

	for i := 0; i < 30; i++ {
		wg.Add(1)
		cashierIndex = cashierIndex % 3
		func(cashier func(int), i int) {
			// Making an order
			go cashier(i)
		}(cashiers[cashierIndex], i)
		cashierIndex++
	}
	wg.Wait()

}
