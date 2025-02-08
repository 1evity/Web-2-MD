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

	// Create a new goja vm
	vm := goja.New()

	// Insert the HTML content into the virtual machine as a global scope variable
	if err := vm.Set("htmlContent", string(htmlContent)); err != nil {
		return "", fmt.Println("failed to set HTML content into VM: %v", err)
	}

	// Define the JavaScript code
	script := `
		const { JSDOM } = require('jsdom');
        const dom = new JSDOM(htmlContent);
        const reader = new Readability(dom.window.document);
        const article = reader.parse();
        article.content;
	`

	// Execute the JS code
	result, err := vm.RunString(script)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML with readability.js: %v", err)
	}

	// Return the decluttered HTML content
	declutteredHTML := result.String()
	return declutteredHTML, nil
}

