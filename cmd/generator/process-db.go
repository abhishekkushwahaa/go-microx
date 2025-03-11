package generator

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/fatih/color"
)

// Function to process database template before writing
func processDatabaseTemplate(templateFile, outputFile, projectName, database string) error {
	tplContent, err := os.ReadFile(templateFile)
	if err != nil {
		color.Red("❌ Failed to read template file: %v", err)
		return err
	}

	// Define the template data
	data := struct {
		ProjectName string
		Database    string
	}{
		ProjectName: projectName,
		Database:    database,
	}

	// Parse and execute the template
	tpl, err := template.New("database").Parse(string(tplContent))
	if err != nil {
		color.Red("❌ Failed to parse template: %v", err)
		return err
	}

	var output bytes.Buffer
	err = tpl.Execute(&output, data)
	if err != nil {
		color.Red("❌ Failed to execute template: %v", err)
		return err
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(outputFile), os.ModePerm); err != nil {
		color.Red("❌ Failed to create directories: %v", err)
		return err
	}

	// Write the processed template
	err = os.WriteFile(outputFile, output.Bytes(), 0644)
	if err != nil {
		color.Red("❌ Failed to write output file: %v", err)
		return err
	}
	return nil
}
