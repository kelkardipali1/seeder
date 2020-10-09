package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func readJson(filePath string) ([]map[string]interface{}, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	ab, mb := new(bytes.Buffer), new(bytes.Buffer)
	mw := io.MultiWriter(ab, mb)

	_, err = mw.Write(b)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	ad := json.NewDecoder(ab)
	var data []map[string]interface{}
	err = ad.Decode(&data)
	if err == nil {
		return data, nil
	}

	md := json.NewDecoder(mb)
	var mData map[string]interface{}
	err = md.Decode(&mData)
	if err == nil {
		return []map[string]interface{}{mData}, nil
	}
	log.Println(err.Error())
	return nil, err
}
