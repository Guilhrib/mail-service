package mailservice

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"roboteasy.com/main-service/internal/models/dtos"
)

const (
	clietnId     = "e10e427c-dc8c-4d03-b11e-67381a3376e1"
	tenantId     = "97741127-d91c-4a45-8df5-6e6743e95aa1"
	clientSecret = "voJ8Q~Q0o4piI64nlyVNzxaq-zS1~~GPSR6LGb3h"
	username     = "sistema@agapys.com"
	password     = "Mub66180"
	requestUrl   = "https://login.microsoftonline.com/" + tenantId + "/oauth2/v2.0/token"
)

func GetToken() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", clietnId)
	data.Set("client_secret", clientSecret)
	data.Set("username", username)
	data.Set("password", password)
	data.Set("scope", "https://graph.microsoft.com/.default")

	payload := bytes.NewBufferString(data.Encode())

	req, err := http.NewRequest(http.MethodPost, requestUrl, payload)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return "", err
	}

	response := dtos.TokenDTO{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	tokenResponse := response.TokenType + " " + response.AccessToken
	return tokenResponse, nil
}
