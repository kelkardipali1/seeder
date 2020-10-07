package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)



func readJson(filePath string) interface{} {
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

	var data map[string]interface{}
	err = ad.Decode(&data)
	if err == nil {
		fmt.Println(1)
		return data
	}

	md := json.NewDecoder(mb)

	var mdata []map[string]interface{}
	err = md.Decode(&mdata)
	if err == nil {
		fmt.Println("mdata",mdata)
		for _, j := range mdata {
			fmt.Println("---->",j)
		}
		return mdata
	}
	fmt.Println("unable to read json")

	return nil
}
