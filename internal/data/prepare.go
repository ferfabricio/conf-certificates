package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Attendant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Code  string `json:"code"`
}

type City struct {
	Name       string      `json:"name"`
	Date       string      `json:"date"`
	Attendants []Attendant `json:"attendants"`
}

type CertificateData struct {
	Cities []City `json:"cities"`
}

// GetCollection return collection of certificates to be generated
func GetCollection(file string) (*CertificateData, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cData := &CertificateData{}
	err = json.Unmarshal([]byte(data), &cData)
	if err != nil {
		return nil, err
	}

	return cData, nil
}

func Print(cData *CertificateData) {
	for i := 0; i < len(cData.Cities); i++ {
		fmt.Println("City:", cData.Cities[i].Name)
		fmt.Println("Date:", cData.Cities[i].Date)
		for j := 0; j < len(cData.Cities[i].Attendants); j++ {
			a := cData.Cities[i].Attendants[j]
			fmt.Println(a.Name, " - ", a.Email, " - ", a.Code)
		}
	}
}
