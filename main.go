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
	client := NewHTTPRequest()
	client.createRequest(url, data)

}
