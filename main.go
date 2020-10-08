package main

import (
	"flag"
)

func main() {
	var url, filePath string
	flag.StringVar(&url, "url", "", "endpoint url")
	flag.StringVar(&filePath, "filePath", "", "json file path ")
	flag.Parse()

	data := readJson(filePath)
	client := NewHttpClient()
	client.doSomething(url, data)

}
