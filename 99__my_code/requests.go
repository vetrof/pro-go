package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	// Чтение тела ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(body)) // Вывод тела ответа
}
