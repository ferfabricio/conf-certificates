package main

import (
	"os"

	"github.com/ferfabricio/certificados-devparana-go/internal/data"
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
}
