package main

import (
	"fmt"
	"money"
)

func main() {

	a := money.NewMoney(1000, -2) // Corrected for clarity: 166004 * 10^-3 = 166.004

	b, _ := money.NewFromString("1000")

	c, _ := b.Divide(a, 2)

	fmt.Println("Result of division (c):", c.String())
}
