package postcalls

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type requestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PostCall(postUrl string, username string, password string) (*http.Response, error) {

	data := &requestBody{
		username,
		password,
	}

	reqBody, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	//TODO SET CONTECT TYPE application/json

	request, err := http.NewRequest("POST",
		postUrl,
		strings.NewReader(string(reqBody)))

	if err != nil {
		log.Fatal("Error On Crafting Reuqest", err)
	}

	request.Header.Set("Content-Type", "application/json")

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
