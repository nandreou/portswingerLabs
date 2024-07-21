package postcalls

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func PostCall(postUrl string, username string, password string) (*http.Response, error) {

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

func PasswordChange(postUrl *string, username string, password string, newPass string, newPassConfirm string, cookie *string) (*http.Response, error) {
	data := &url.Values{}

	data.Add("username", username)
	data.Add("current-password", password)
	data.Add("new-password-1", newPass)
	data.Add("new-password-2", newPassConfirm)

	request, err := http.NewRequest("POST",
		*postUrl,
		strings.NewReader(data.Encode()))

	if err != nil {
		log.Fatal("Error On Crafting Reuqest", err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(&http.Cookie{Name: "session", Value: *cookie})

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
