//Filename:main.go
//Waitgroups:Another mechanism of concurrency and synchronization
//Keeps track of goroutines and makes main() to wait until they are executed
//Does not require time.sleep()

package main

import (
	"fmt"
	"sync"
)

func festival() {

	fmt.Println("Christmas is the best annual christian festival")
}

func wish() {

	fmt.Println("I wish a had a mansion")
}

// main is a goroutine
func main() {
	//create a waitgroup
	var wg sync.WaitGroup
	//increment waitgroup by 1
	wg.Add(1)
	//anonymous goroutine calls festival()
	go func() {
		festival()
		wg.Done() //decrements wg by 1
	}()
	//increment waitgroup 1
	wg.Add(1)

	//anonymous goroutine calls wish()
	go func() {
		wish()
		wg.Done() //decrements wg by 1
	}()
	//makes main()to wait for wg to become 0
	wg.Wait()

	fmt.Println("Happy Birthday to you!")

}
