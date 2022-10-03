package main

import (
	"os"

	"github.com/ferfabricio/certificados-devparana-go/internal/data"
	"github.com/ferfabricio/certificados-devparana-go/internal/pdf"
)

func main() {
	a := os.Args
	if len(a) == 1 {
		panic("missing required argument")
	}

	c, err := data.GetCollection(a[1])
	if err != nil {
		panic(err)
	}

	data.Print(c)

	for i := 0; i < len(c.Cities); i++ {
		for j := 0; j < len(c.Cities[i].Attendants); j++ {
			err = pdf.SavePDF(&c.Cities[i], c.Cities[i].Attendants[j])
			if err != nil {
				panic(err)
			}
		}
	}
}
