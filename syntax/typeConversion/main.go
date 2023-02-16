package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100"

	//String to int
	intVar, _ := strconv.Atoi(str)
	fmt.Println(intVar)

	//String to int64
	int64Var, _ := strconv.ParseInt(str, 10, 64) //strconv.ParseInt(s, 10, 64)	for int64
	fmt.Println(int64Var)

	//int to string
	intstr := strconv.Itoa(intVar)
	fmt.Println(intstr)

	//int64 to string
	int64str := strconv.FormatInt(int64Var, 10)
	fmt.Println(int64str)
}
