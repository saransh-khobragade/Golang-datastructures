package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func convertStringToArrayOfInt(str string) []int {

	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")

	strArr := strings.Split(str, ",")

	intArr := make([]int, len(strArr))
	for i := range intArr {
		intArr[i], _ = strconv.Atoi(strArr[i])
	}
	return intArr
}

func readFile(filepath string) ([]int, []int) {

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	result := string(content)
	inputLineArr := strings.Split(result, "\n")

	inputLine1 := inputLineArr[0]
	inputLine2 := inputLineArr[1]

	inputLine1Arr := convertStringToArrayOfInt(inputLine1)
	inputLine2Arr := convertStringToArrayOfInt(inputLine2)

	return inputLine1Arr, inputLine2Arr

}
