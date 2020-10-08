package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readJson(filePath string) []map[string]interface{} {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	b, _ := ioutil.ReadAll(f)

	ab, mb := new(bytes.Buffer), new(bytes.Buffer)
	mw := io.MultiWriter(ab, mb)

	_, err = mw.Write(b)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	ad := json.NewDecoder(ab)
	var data []map[string]interface{}
	err = ad.Decode(&data)
	if err == nil {
		return data
	}


	md := json.NewDecoder(mb)
	var mData map[string]interface{}
	err = md.Decode(&mData)
	if err == nil {
		return []map[string]interface{}{mData}
	}

	return nil
}