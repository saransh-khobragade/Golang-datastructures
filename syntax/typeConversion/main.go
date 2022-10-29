package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	str := "100"

	intVar, _ := strconv.ParseInt(str, 0, 8)
	fmt.Println(intVar, reflect.TypeOf(intVar)) //100 int64

}
