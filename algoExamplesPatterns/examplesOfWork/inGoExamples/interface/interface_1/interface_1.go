package main

import (
	"errors"
	"fmt"
)

type UserWallet struct {
	User  User
	Id    int
	Value uint
}

type User struct {
	Items []int
	Name  string
}

type Item struct {
	Id    int
	Price uint
}

func (uw *UserWallet) Pay(item Item) error {
	if uw.Value >= item.Price {
		uw.Value -= item.Price
		uw.User.Items = append(uw.User.Items, item.Id)
		fmt.Printf("Now your balance is: %v\n", uw.Value)
		return nil
	} else {
		return errors.New("Not enough value on account to buy this product")
	}
}

type Payer interface {
	Pay(Item) error
}

func Buy(p Payer, item Item) {
	err := p.Pay(item)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Thank you for purchasing!\n")
}

func main() {
	u := User{Name: "Jojo", Items: make([]int, 0)}
	uw := &UserWallet{User: u, Id: 1, Value: 150}
	i := Item{Id: 17, Price: 50}

	Buy(uw, i)
}
