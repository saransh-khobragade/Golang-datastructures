package main

import (
	"fmt"
	"reflect"
)

func main() {

	//INTEGER only intialization with type
	var i int
	//int64 -9223372036854775808 til **9223372036854775807.
	//int -2147483648 and 2147483647

	//FLOAT intialization and declartion  with type
	var j float32 = 20.8
	j = 30.8

	//STRING intialization and declartion  with type
	var k = "20"
	k = "30"

	//BOOLEAN no var,no type,shortcut
	l := true

	//MULTIPLE declaration
	m, n := 10, 20

	//CONST var
	const o string = "abc"

	fmt.Println(i, reflect.TypeOf(i))                       //0 int
	fmt.Println(j, reflect.TypeOf(j))                       //30.8 float32
	fmt.Println(k, reflect.TypeOf(k))                       //30 string
	fmt.Println(l, reflect.TypeOf(l))                       //true bool
	fmt.Println(m, n, reflect.TypeOf(m), reflect.TypeOf(m)) //10 20 int int
	fmt.Println(o, reflect.TypeOf(o))                       //abc string
}
