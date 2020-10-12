package main

import (
	"flag"
	"log"
)

func main() {
	var url, filePath string
	flag.StringVar(&url, "url", "", "endpoint url")
	flag.StringVar(&filePath, "filePath", "", "json file path ")
	flag.Parse()

	data, err := readJson(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	client := NewHttpClient()
	req := NewHTTPRequest(client)

	error := req.createRequest(url, data)
	if len(error) != 0 {
		for _, err := range error {
			log.Println(err.Error())
		}
		return
	}

	log.Print("Done")
}
