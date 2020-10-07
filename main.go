package main

import "fmt"

func main() {
	data := readJson("./seed.json")

	fmt.Println(data)
}
