package declutter

import {
	"fmt"
	"io/ioutil"
	"os/exec"
	"github.com/dop251/goja"
}

// DeclutterHTML takes a HTML file, processes it with readability.js, and returns the decluttered HTML.
func DeclutterHTML(htmlFile string) (string, error) {
	// Read the HTML File
	fmt.Println("reading HTML file ...")
	htmlContent, err := ioutil.ReadFile(htmlFile)
	if err != nil {
		return "", fmt.Errorf("failed to read HTML file: %v", err)
	}

	// Read the readability.js script
	readability, err := ioutil.ReadFile("./scripts/readability/readability.js")
	fmt.Println("importing readability.js library ...")
	if err != nil {
		return "", fmt.Errorf("failed to import readability.js: %v", err)
	}

	
}

