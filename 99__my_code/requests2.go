package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "https://dog.ceo/api/breeds/image/random"
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		fmt.Println(resp.Status)
	} else {
		fmt.Fprintf(os.Stderr, "get error: %v\n", err)
		os.Exit(1)
	}

}
