package utils

import (
	"log"
	"os"
)

func ReadFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Panicf("read %s files failed. err: %v", filename, err)
	}
	return data
}
