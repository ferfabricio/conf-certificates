package pdf

import (
	"fmt"

	"github.com/ferfabricio/certificados-devparana-go/internal/data"
	"github.com/ferfabricio/certificados-devparana-go/internal/pdf/templates"
)

// SavePDF generate and save the PDF file
func SavePDF(city *data.City, attendent data.Attendant) error {
	var err error
	pdf, err := templates.Prepare(attendent, city)
	if err != nil {
		return err
	}

	err = pdf.WritePdf(fmt.Sprintf("./output/%s.pdf", attendent.Code))
	if err != nil {
		return err
	}

	return nil
}
