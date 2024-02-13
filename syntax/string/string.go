package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {

	arr := "ABCD"
	for i := range arr {
		fmt.Println(arr[i])         //65 66 67 68
		fmt.Println(string(arr[i])) //A B C D
	}

	for i := range arr {
		if arr[i] == 'B' {
			fmt.Println("B is there")
		}
	}

	//Split
	fmt.Println(strings.Split("hello,my,name,is,saransh", ","))

	//Join
	t := []string{"hello", "my", "name"}
	fmt.Println(strings.Join(t, ","))

	//Lowercase
	s := "ABC"
	str := strings.ToLower(s)
	fmt.Println(str)

	//Uppercase
	str = strings.ToUpper(s)
	fmt.Println(str)

	//Sort string
	m := SortString("dcba")
	fmt.Println(m)

	//String to int
	fmt.Println(strconv.Atoi("5"))

	//Int to string
	fmt.Println(strconv.Itoa(15))

}
