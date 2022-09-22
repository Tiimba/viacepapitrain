package main

import (
	cepai "apicep/src"
	"flag"
	"os"
)

func main() {
	var cep string
	flag.StringVar(&cep, "cep", "13099160", "CEP necessário para incluir")
	flag.Parse()
	cepai.GetRequest(os.Getenv("URL") + string(cep) + "/json/")
}
