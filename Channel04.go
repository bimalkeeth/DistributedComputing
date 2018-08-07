package main

import (
	"fmt"
	"sync"
)

func cashier(chasierID int, orderChannel <-chan int, wg *sync.WaitGroup) {
	for orderprocessed := 0; orderprocessed < 10; orderprocessed++ {
		orderNum := <-orderChannel
		fmt.Println("Cashier ", chasierID, "Processing order", orderNum,
			"Orders Processed", orderprocessed)
		wg.Done()
	}
}
func main() {

	var wg sync.WaitGroup
	wg.Add(30)
	ordersChannel := make(chan int)
	for i := 0; i < 3; i++ {
		// Start the three cashiers
		func(i int) {
			go cashier(i, ordersChannel, &wg)
		}(i)
	}
	// Start adding orders to be processed.
	for i := 0; i < 30; i++ {
		ordersChannel <- i
	}
	wg.Wait()

}
