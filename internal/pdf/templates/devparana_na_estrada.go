package templates

import (
	"github.com/ferfabricio/certificados-devparana-go/internal/data"
	"github.com/signintech/gopdf"
)

func DevParanaNaEstrada2022(attendent data.Attendant, city *data.City) (*gopdf.GoPdf, error) {
	var err error

	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape})
	pdf.AddPage()

	tpl := pdf.ImportPage("./assets/pdf/template-devparana-na-estrada-2022.pdf", 1, "/MediaBox")
	pdf.UseImportedTemplate(tpl, 0, 0, gopdf.PageSizeA4Landscape.W, gopdf.PageSizeA4Landscape.H)

	err = pdf.AddTTFFont("aleo", "./assets/fonts/Aleo-Bold.ttf")
	if err != nil {
		return nil, err
	}
	err = pdf.AddTTFFont("opensans", "./assets/fonts/Open-Sans.ttf")
	if err != nil {
		return nil, err
	}

	err = pdf.SetFont("aleo", "", 25)
	if err != nil {
		return nil, err
	}

	s, err := pdf.MeasureTextWidth(attendent.Name)
	if err != nil {
		return nil, err
	}

	pdf.SetXY((gopdf.PageSizeA4Landscape.W/2)-(s/2), 270)

	err = pdf.Text(attendent.Name)
	if err != nil {
		return nil, err
	}

	err = pdf.SetFont("opensans", "", 13)
	if err != nil {
		return nil, err
	}

	pdf.SetXY(350, 400)

	err = pdf.Text(city.Name)
	if err != nil {
		return nil, err
	}

	pdf.SetXY(350, 420)

	err = pdf.Text(city.Date)
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
