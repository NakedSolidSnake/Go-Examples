package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account wht balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account wht balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Mutex Tutorial")

	balance = 1000

	var wg sync.WaitGroup
	wg.Add(2)

	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)

}
