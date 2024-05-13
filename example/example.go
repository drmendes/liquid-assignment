package main

import (
	"fmt"
	"money"
)

func main() {

	a := money.NewMoney(217, -2)
	b := money.NewMoney(4, 0)

	c := b.Add(a)
	fmt.Println("Result:", c.String())

}
