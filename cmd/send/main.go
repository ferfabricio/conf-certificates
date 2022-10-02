package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	m := mail.NewV3Mail()

	from := mail.NewEmail("DevParaná", "contato@devparana.org")
	to := mail.NewEmail("Everton Tavares", "ezidiu@gmail.com")

	m.SetFrom(from)
	m.SetTemplateID(os.Getenv("SENDGRID_CERTIFICATE_TEMPLATE"))

	p := mail.NewPersonalization()
	p.AddTos(to)
	p.SetDynamicTemplateData("first_name", "Everton Tavares")

	m.AddPersonalizations(p)

	a := mail.NewAttachment()
	c, err := os.ReadFile("./output/83064af3-bb81-4514-a6d4-afba340825cd.pdf")
	if err != nil {
		panic(err)
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(c))

	a.SetContent(encoded)
	a.SetType("application/pdf")
	a.SetFilename("certificado.pdf")
	a.SetDisposition("attachment")

	m.AddAttachment(a)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

}
