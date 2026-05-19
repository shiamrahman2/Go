package main

import "fmt"

type People interface {
	PrintDetails()
	ReceiveAmount(amount float64) float64
}
type BankDetails interface {
	WithDrawMoney(amount float64) float64
}

type user struct {
	Name  string
	Age   int
	Money float64
}

func (usr user) PrintDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
	fmt.Println("Money:", usr.Money)
}
func (usr user) ReceiveAmount(amount float64) float64 {
	usr.Money = usr.Money + amount
	return usr.Money
}
func (usr user) WithDrawMoney(amount float64) float64 {
	usr.Money = usr.Money - amount
	return usr.Money
}
func main() {
	usr1 := user{
		Name:  "Shiam Hosen Monna",
		Age:   22,
		Money: 900.00,
	}
	usr1.PrintDetails()
	mn := usr1.ReceiveAmount(500)
	fmt.Println("Receive Money", mn)
	mn = usr1.WithDrawMoney(300)
	fmt.Println("Money After Withdraw:", mn)
}
