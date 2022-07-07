package ExercisesDay2

import (
	"fmt"
	"sync"
)

type BankAccounts struct {
	mu             sync.Mutex
	accountDetails map[string]int
}

func (ba *BankAccounts) deposit(accountNumber string, amount int) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	initialAmount := ba.accountDetails[accountNumber]
	finalAmount := initialAmount + amount
	ba.accountDetails[accountNumber] = finalAmount
	fmt.Println("initial balance:", initialAmount, ", amount deposited:", amount, ", current balance", finalAmount)
}
func (ba *BankAccounts) withdraw(accountNumber string, amount int) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	initialAmount := ba.accountDetails[accountNumber]
	finalAmount := initialAmount
	if initialAmount-amount < 0 {
		//panic("can't withdraw an amount more than current balance")
		fmt.Println("Can't withdraw an amount more than current balance, ")
	} else {
		finalAmount -= amount
		ba.accountDetails[accountNumber] = finalAmount
	}
	fmt.Println("initial balance:", initialAmount, ", amount deposited:", amount, ", current balance", finalAmount)
}

func Problem3() {
	accounts := BankAccounts{
		accountDetails: map[string]int{"acc1": 500},
	}
	var wg sync.WaitGroup
	withdraw := func(accountNumber string, amount int) {
		accounts.withdraw(accountNumber, amount)
		wg.Done()
	}
	deposit := func(accountNumber string, amount int) {
		accounts.deposit(accountNumber, amount)
		wg.Done()
	}
	wg.Add(12)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	wg.Wait()

	fmt.Println(accounts.accountDetails["acc1"])
}
