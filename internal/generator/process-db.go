package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/fatih/color"
)

func processDatabaseTemplate(templateFile, outputFile, projectName, database string) error {
	tmplContent, err := os.ReadFile(templateFile)
	if err != nil {
		color.Red("Failed to read template file: %v", err)
		return err
	}

	data := struct {
		ProjectName string
		Database    string
	}{
		ProjectName: projectName,
		Database:    database,
	}

	tmpl, err := template.New("database").Parse(string(tmplContent))
	if err != nil {
		color.Red("Failed to parse template: %v", err)
		return err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, data)
	if err != nil {
		color.Red("Failed to execute template: %v", err)
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outputFile), os.ModePerm); err != nil {
		color.Red("Failed to create directories: %v", err)
		return err
	}

	err = os.WriteFile(outputFile, output.Bytes(), 0644)
	if err != nil {
		color.Red("Failed to write output file: %v", err)
		return err
	}
	return nil
}

func getDatabaseImport(serviceName, database string) string {
	if database == "None" {
		return ""
	}
	return fmt.Sprintf(`"%s/database"`, serviceName)
}
