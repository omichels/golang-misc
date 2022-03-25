package main

import (
	"fmt"
	"golang-misc/singleton"
	"sync"
	"time"
)

var waitGroupd sync.WaitGroup

func main() {

	waitGroupd.Add(2)

	go dog()
	go cat()

	waitGroupd.Wait()

}

func dog() {
	for i := 1; i <= 5; i++ {
		time.Sleep(2 * time.Duration(time.Millisecond))
		fmt.Printf("dog %d\n", singleton.GetSingletonInstance().GetNextId())
	}
	waitGroupd.Done()
}

func cat() {
	for i := 1; i <= 5; i++ {
		time.Sleep(3 * time.Duration(time.Millisecond))
		fmt.Printf("cat %d\n", singleton.GetSingletonInstance().GetNextId())
	}
	waitGroupd.Done()
}
