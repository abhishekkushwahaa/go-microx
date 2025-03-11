package generator

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func getMicroservices(projectType string) []string {
	switch projectType {
	case "E-commerce":
		return []string{"user-service", "product-service", "order-service", "payment-service", "cart-service"}
	case "Video-Streaming":
		return []string{"user-service", "video-service", "subscription-service", "recommendation-service", "analytics-service"}
	case "Food-Delivery":
		return []string{"user-service", "restaurant-service", "order-service", "delivery-service", "payment-service"}
	default:
		return []string{"api-service"}
	}
}

// GenerateProject copies the selected template to the target directory
func GenerateProject(projectType, projectName, database, httpRouter, authMethod string) {
	templatePath := filepath.Join("templates", projectType)
	targetPath := filepath.Join(".", projectName)

	color.Cyan("📂 Creating project: %s", projectName)
	color.Green("📦 Type: %s", projectType)
	color.Yellow("🫙  Database: %s", database)
	color.Blue("🌐 HTTP Router: %s", httpRouter)
	color.White("🔑 Auth Method: %s", authMethod)

	// Copy template files to target project
	err := copyDirectory(templatePath, targetPath)
	if err != nil {
		color.Red("❌ Failed to create project: %v", err)
		return
	}

	// Fetch microservices dynamically
	microservices := getMicroservices(projectType)

	if database != "None" {
		for _, service := range microservices {
			servicePath := filepath.Join(targetPath, service)
			dbTemplate := filepath.Join("cmd", "built", "database.tmpl")
			dbTarget := filepath.Join(servicePath, "database", "db.go")

			if err := os.MkdirAll(filepath.Dir(dbTarget), os.ModePerm); err != nil {
				color.Red("❌ Failed to create directories for %s: %v", service, err)
				return
			}
			processDatabaseTemplate(dbTemplate, dbTarget, service, database)
		}
	}

	replacePlaceholders(targetPath, projectName, database, httpRouter, authMethod)

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
func replacePlaceholders(rootPath, projectName, database, httpRouter, authMethod string) {
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
		updatedContent = strings.ReplaceAll(updatedContent, "{{HTTP_ROUTER}}", httpRouter)
		updatedContent = strings.ReplaceAll(updatedContent, "{{AUTH_METHOD}}", authMethod)

		segments := strings.Split(path, string(os.PathSeparator))
		if len(segments) > 1 {
			serviceName := segments[len(segments)-2]
			updatedContent = strings.ReplaceAll(updatedContent, "{{SERVICE_NAME}}", serviceName)
		}

		return ioutil.WriteFile(path, []byte(updatedContent), info.Mode())
	})
}
