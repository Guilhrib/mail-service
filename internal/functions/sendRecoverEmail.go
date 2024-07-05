package functions

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	mailservice "roboteasy.com/main-service/internal"
	"roboteasy.com/main-service/internal/models/dtos"
	"roboteasy.com/main-service/internal/models/enums"
)

const (
	requestUrl = "https://graph.microsoft.com/v1.0/me/sendMail"
)

func SendRecoverEmail(email string) error {
	token, err := mailservice.GetToken()
	if err != nil {
		return err
	}

	payload := buildNewEmail()

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 202 {
		log.Println(res)
		return err
	}

	log.Println("Email sent")
	return nil
}

func buildNewEmail() dtos.EmailDTO {
	return dtos.EmailDTO{
		Message: dtos.MessageDTO{
			Subject: "Recuperação de senha",
			Body: dtos.BodyDTO{
				ContentType: enums.HTML,
				Content:     buildEmailTemplate().String(),
			},
			ToRecipients: []dtos.RecipientDTO{
				{
					EmailAddress: dtos.AddressDTO{
						Address: "guilherme.ribeiro@agapys.com",
					},
				},
			},
			Attachments: []dtos.AttachmentDTO{},
		},
	}
}

func buildEmailTemplate() *bytes.Buffer {
	t, _ := template.ParseFiles("internal/templates/template-test.html")

	var body bytes.Buffer

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Guilherme Ribeiro BLABLA",
		Message: "This is a test message in a HTML template",
	})

	return &body
}
