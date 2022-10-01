package main

import "github.com/ferfabricio/certificados-devparana-go/internal/data"

func main() {
	c, err := data.GetCollection("../assets/json/example.json")
	if err != nil {
		panic(err)
	}

	data.Print(c)
}
