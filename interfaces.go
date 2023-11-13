package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(0)
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(body))
}
