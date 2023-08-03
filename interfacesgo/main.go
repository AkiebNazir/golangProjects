package main

import (
	"fmt"
)

type Person struct {
	name    string
	address string
}

func (p *Person) PrintData() {
	fmt.Println(p.name, p.address)
}
func main() {
	p := Person{
		name:    "Akieb",
		address: "Nagbale",
	}

	p.PrintData()

	// var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan string)
	for i := 0; i < 5; i++ {
		// wg.Add(1)
		go func() {
			ch1 <- "hello World"
			// wg.Done()
		}()
		// wg.Add(1)
		go func() {
			ch2 <- "ending ... !"
			// wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}

	// wg.Wait()
	fmt.Println("Process completed")

}
