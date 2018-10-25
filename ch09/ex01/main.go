package bank

import (
	"fmt"
)

type withdraw struct {
	amount   int
	resultCh chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdraw)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	r := make(chan bool)
	withdraws <- withdraw{amount, r}
	return <-r
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdraws:
			fmt.Println(balance)
			if balance < w.amount {
				w.resultCh <- false
			} else {
				balance -= w.amount
				w.resultCh <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
