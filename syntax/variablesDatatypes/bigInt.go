package main

import (
	"fmt"
	"math/big"
)

func main() {

	//BigInt Bigger than int64 9223372036854775807

	i := big.NewInt(13)
	fmt.Println(i)

	//Converting value more than int64 to bigint
	j := new(big.Int)
	j.SetString("9223372036854775808", 10)
	fmt.Println(j)

	//Adding to bigInt
	j.Add(j, i)
	fmt.Println("ADD", j)

	//Subtractig from BigInt
	j.Sub(j, i)
	fmt.Println("SUB", j)

	//Divide from BigInt
	j.Div(j, i)
	fmt.Println("DIVIDE", j)

	//Modulus of BigInt
	mod := new(big.Int)
	mod.Mod(j, i)
	fmt.Println("Mod", mod)

	//Converting bigInt to int
	fmt.Println("bigInt to int", j.Int64())

}
