package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	fmt.Println(today)          //2022-07-30 13:57:24.53317 +0530 IST m=+0.000398005
	fmt.Println(today.Day())    //30
	fmt.Println(today.Month())  //July
	fmt.Println(today.Year())   //2022
	fmt.Println(today.Hour())   //13
	fmt.Println(today.Minute()) //57
	fmt.Println(today.Second()) //24
}
