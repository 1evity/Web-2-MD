package save

import (
	"fmt"
	"io/ioutil"
)

// SaveFile takes in content and writes it to a new file
func SaveFile(inputContent, outputFile string) error {
	// Write output file
	err := ioutil.WriteFile(outputFile, []byte(inputFile), 0644)

	// Check for error
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}