package utils

import (
	"io/ioutil"
	"log"
)

func ReadHttpFile(filepath string) (file []byte, err error) {
	file, e := ioutil.ReadFile(filepath)
	if e != nil {
		log.Println("Error while trying to read http file at", filepath, "error:", e)
		return nil, e
	}

	return file, nil
}
