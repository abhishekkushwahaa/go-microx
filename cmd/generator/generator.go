package generator

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// GenerateProject copies the selected template to the target directory
func GenerateProject(projectType, projectName, database string) {
	templatePath := filepath.Join("templates", projectType)
	targetPath := filepath.Join(".", projectName)

	color.Cyan("📂 Creating project: %s", projectName)
	color.Green("📦 Type: %s", projectType)
	color.Yellow("🫙  Database: %s", database)

	// Copy template files to target project
	err := copyDirectory(templatePath, targetPath)
	if err != nil {
		color.Red("❌ Failed to create project: %v", err)
		return
	}

	// Replace placeholders inside the project
	replacePlaceholders(targetPath, projectName, database)

	color.Magenta("✅ Project %s has been successfully created!\n", projectName)
}

// Copy all files from template to target directory
func copyDirectory(src, dst string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(src, path)
		targetPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(targetPath, content, info.Mode())
	})
}

// Replace placeholders in files
func replacePlaceholders(rootPath, projectName, database string) {
	filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		updatedContent := strings.ReplaceAll(string(content), "{{PROJECT_NAME}}", projectName)
		updatedContent = strings.ReplaceAll(updatedContent, "{{DATABASE}}", database)

		segments := strings.Split(path, string(os.PathSeparator))
		if len(segments) > 1 {
			serviceName := segments[len(segments)-2]
			updatedContent = strings.ReplaceAll(updatedContent, "{{SERVICE_NAME}}", serviceName)
		}

		return ioutil.WriteFile(path, []byte(updatedContent), info.Mode())
	})
}
