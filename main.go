package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func GetRequest(url string) Address {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var address Address
	json.Unmarshal(body, &address)
	fmt.Println(address)
	return address
}

// func main() {
// 	response := GetRequest("https://viacep.com.br/ws/" + os.Getenv("CEP") + "/json/")
// 	fmt.Println(re)

// }

// function which return "geeks"
func ReturnGeeks() string {
	return "geeks"
}
