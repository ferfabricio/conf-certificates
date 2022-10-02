package pdf

import (
	"fmt"

	"github.com/ferfabricio/certificados-devparana-go/internal/data"
	"github.com/signintech/gopdf"
)

// SavePDF generate and save the PDF file
func SavePDF(city string, date string, attendent data.Attendant) error {
	var err error
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape})
	pdf.AddPage()

	tpl := pdf.ImportPage("./assets/pdf/template.pdf", 1, "/MediaBox")
	pdf.UseImportedTemplate(tpl, 0, 0, gopdf.PageSizeA4Landscape.W, gopdf.PageSizeA4Landscape.H)

	err = pdf.AddTTFFont("aleo", "./assets/fonts/Aleo-Bold.ttf")
	if err != nil {
		return err
	}
	err = pdf.AddTTFFont("opensans", "./assets/fonts/Open-Sans.ttf")
	if err != nil {
		return err
	}

	err = pdf.SetFont("aleo", "", 25)
	if err != nil {
		return err
	}

	s, err := pdf.MeasureTextWidth(attendent.Name)
	if err != nil {
		return err
	}

	pdf.SetXY((gopdf.PageSizeA4Landscape.W/2)-(s/2), 270)

	pdf.Text(attendent.Name)

	pdf.SetFont("opensans", "", 13)

	pdf.SetXY(350, 400)

	pdf.Text(city)

	pdf.SetXY(350, 420)

	pdf.Text(date)

	pdf.WritePdf(fmt.Sprintf("./output/%s.pdf", attendent.Code))

	return nil
}
