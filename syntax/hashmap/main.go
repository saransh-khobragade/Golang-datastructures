package main

import "fmt"

func main() {
	var obj = map[string]string{
		"a": "1",
		"b": "2",
	}

	obj2 := make(map[string]string)
	obj2["1"] = "a"
	obj2["2"] = "b"
	obj2["3"] = "c"

	fmt.Println(obj)  //map[a:1 b:2]
	fmt.Println(obj2) //map[1:a 2:b 3:c]

	delete(obj2, "2")

	fmt.Println(obj2) //map[1:a 3:c]

	for i, v := range obj2 {
		println(i, v)
	}

	//check key exists in map
	if val, ok := obj["key"]; ok {
		fmt.Println(val)
	}

}
