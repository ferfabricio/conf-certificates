package templates

import (
	"errors"

	"github.com/ferfabricio/certificados-devparana-go/internal/data"
	"github.com/signintech/gopdf"
)

func getTemplate(t string, a data.Attendant, c *data.City, p *gopdf.GoPdf) (*gopdf.GoPdf, error) {
	switch t {
	case "devparana_na_estrada_2022":
		return DevParanaNaEstrada2022(a, c)
	}
	return nil, errors.New("template not implemented")
}

func Prepare(a data.Attendant, c *data.City) (*gopdf.GoPdf, error) {
	pdf := &gopdf.GoPdf{}

	return getTemplate(c.Template, a, c, pdf)
}
