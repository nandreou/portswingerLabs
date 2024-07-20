package httpcalls

import "net/http"

func LogIn(url string, hash string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.AddCookie(&http.Cookie{Name: "stay-logged-in", Value: hash})
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
