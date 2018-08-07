package main

import "fmt"

func main() {
	fcn1 := func(i int) {
		fmt.Println("fcn1", i)
	}
	fcn2 := func(i int) {
		fmt.Println("fcn2", i)
	}
	fcn3 := func(i int) {
		fmt.Println("fcn3", i)
	}

	ch := make(chan func(int))
	done := make(chan bool)

	go func() {
		for fcn := range ch {
			fcn(10)
		}
		fmt.Println("Exiting")
		done <- true
	}()

	ch <- fcn1
	ch <- fcn2
	ch <- fcn3

	close(ch)
	<-done
	defer close(done)
}
