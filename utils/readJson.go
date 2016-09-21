package utils

import (
	"io/ioutil"
	"log"
)

func ReadJsonFromFile(filepath string) (file []byte, err error) {
	log.Println("Read from file:", filepath)
	// try to read file
	file, e := ioutil.ReadFile(filepath)
	if e != nil {
		log.Println("Error during file read", e)
		return nil, e
	}

	return file, nil
}
