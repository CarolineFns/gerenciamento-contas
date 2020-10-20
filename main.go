package main

import "fmt"

//ContaCorrente is a struct
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaCarol := ContaCorrente{titular: "Carol", saldo: 125}
	contaRafael := ContaCorrente{"Rafael", 5874, 88997, 1000.60}

	fmt.Println(contaCarol)
	fmt.Println(contaRafael)
}
