package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWiter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWiter{}

	io.Copy(lw, resp.Body)

}

func (logWiter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs))
	return len(bs), nil
}
