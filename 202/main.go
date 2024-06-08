//goroutine

package main

import (
	"fmt"
	"sync"
)

func main() {
	bank := &BankAccount{balance: 100.00}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		bank.Deposit(100)
		fmt.Println("Deposit 100 Rs")

	}()

	go func() {
		defer wg.Done()
		bank.Withdrawal(20)
		fmt.Println("withdrawal 20")

	}()
	wg.Wait()
	fmt.Println("Final Balance:", bank.Balance())

}

type BankAccount struct {
	balance float32
	sync.Mutex
}

func (b *BankAccount) Deposit(amount float32) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	b.balance += amount
}

func (b *BankAccount) Withdrawal(amount float32) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	if b.balance >= amount {
		b.balance -= amount
	} else {
		fmt.Println("Insufficent Balance...")
	}
}

func (b *BankAccount) Balance() float32 {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.balance
}
