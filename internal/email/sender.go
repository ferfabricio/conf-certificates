package email

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/ferfabricio/certificados-devparana-go/internal/data"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Send(a *data.Attendant) error {

	m := mail.NewV3Mail()

	from := mail.NewEmail(os.Getenv("SENDGRID_SENDER_NAME"), os.Getenv("SENDGRID_SENDER_EMAIL"))
	to := mail.NewEmail(a.Name, a.Email)

	m.SetFrom(from)
	m.SetTemplateID(os.Getenv("SENDGRID_CERTIFICATE_TEMPLATE"))

	p := mail.NewPersonalization()
	p.AddTos(to)
	p.SetDynamicTemplateData("first_name", a.Name)

	m.AddPersonalizations(p)

	at := mail.NewAttachment()
	c, err := os.ReadFile(fmt.Sprintf("./output/%s.pdf", a.Code))
	if err != nil {
		return err
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(c))

	at.SetContent(encoded)
	at.SetType("application/pdf")
	at.SetFilename("certificado.pdf")
	at.SetDisposition("attachment")

	m.AddAttachment(at)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	}

	fmt.Printf("Email enviado para %s, Status Code: %d\n", a.Email, response.StatusCode)

	return nil
}
