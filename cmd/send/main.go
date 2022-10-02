package main

import (
	"os"

	"github.com/ferfabricio/certificados-devparana-go/internal/data"
	"github.com/ferfabricio/certificados-devparana-go/internal/email"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	a := os.Args
	if len(a) == 1 {
		panic("missing required argument")
	}

	c, err := data.GetCollection(a[1])
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(c.Cities); i++ {
		for j := 0; j < len(c.Cities[i].Attendants); j++ {
			err = email.Send(&c.Cities[i].Attendants[j])
			if err != nil {
				panic(err)
			}
		}
	}
}
