package postcalls

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func PostCall(postUrl string, username string, password string, number int) (*http.Response, error) {

	data := &url.Values{}

	data.Add("username", username)
	data.Add("password", password)

	request, err := http.NewRequest("POST",
		postUrl,
		strings.NewReader(data.Encode()))

	if err != nil {
		log.Fatal("Error On Crafting Reuqest", err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("X-Forwarded-For", strconv.Itoa(number))

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	response, err := client.Do(request)

	if err != nil {
		return nil, err

	}

	return response, nil
}
