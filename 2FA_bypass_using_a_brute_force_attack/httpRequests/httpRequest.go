package httprequests

import (
	"net/http"
	"net/url"
	"strings"
)

func RequestGet(urlGet string) (*http.Response, error) {
	response, err := http.Get(urlGet)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func RequestGet2(urlGet string, session string) (*http.Response, error) {
	request, err := http.NewRequest("GET", urlGet, nil)
	if err != nil {
		return nil, err
	}

	request.AddCookie(&http.Cookie{Name: "session", Value: session})

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func RequestPost1(urlPost1 string, csrf string, session string) (*http.Response, error) {

	var (
		data         = &url.Values{}
		sessionCokie = &http.Cookie{Name: "session", Value: session}
	)

	data.Add("username", "carlos")
	data.Add("password", "montoya")
	data.Add("csrf", csrf)

	request, err := http.NewRequest("POST", urlPost1, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(sessionCokie)

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

func RequestPost2(urlPost string, csrf string, session string, mfa string) (*http.Response, error) {
	var (
		data = &url.Values{}
	)

	data.Add("csrf", csrf)
	data.Add("mfa-code", mfa)

	request, err := http.NewRequest("POST", urlPost, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(&http.Cookie{Name: "session", Value: session})

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
