package main

import (
	"fmt"
)

type Conta struct {
	saldo float64
}

func (c *Conta) depositarDezReais() float64 {
	return c.saldo + 10
}

func main() {
	contaTeste := Conta{saldo: 10}

	contaTeste.depositarDezReais()
	contaTeste.depositarDezReais()

	fmt.Println(contaTeste)
}
