package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Println("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		// exit the function on error
		return []byte(""), inputError
	}
	defer inputFile.Close()
	fb, err := ioutil.ReadAll(inputFile)
	return fb, err
}
