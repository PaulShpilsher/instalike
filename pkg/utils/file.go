package utils

import (
	"log"
	"os"
)

//
// File related helper functions
//

// ReadFile is a wrapper function, it panics on error and returns only valid data
func ReadFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Panicf("read %s files failed. err: %v", filename, err)
	}
	return data
}
