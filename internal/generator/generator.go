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

func GenerateProject(projectType, projectName, database, httpRouter, authMethod string) {
	templatePath := filepath.Join("templates", projectType)
	targetPath := filepath.Join(".", projectName)

	color.Cyan("Creating project: %s", projectName)
	color.Green("Type: %s", projectType)
	color.Yellow("Database: %s", database)
	color.Blue("HTTP Router: %s", httpRouter)
	color.White("Auth Method: %s", authMethod)

	err := copyDirectory(templatePath, targetPath)
	if err != nil {
		color.Red("Failed to create project: %v", err)
		return
	}

	microservices := getMicroservices(projectType)

	if database != "None" {
		for _, service := range microservices {
			servicePath := filepath.Join(targetPath, service)
			dbTemplate := filepath.Join("internal", "templates", "database.tmpl")
			dbTarget := filepath.Join(servicePath, "database", "db.go")

			if err := os.MkdirAll(filepath.Dir(dbTarget), os.ModePerm); err != nil {
				color.Red("Failed to create directories for %s: %v", service, err)
				return
			}
			processDatabaseTemplate(dbTemplate, dbTarget, service, database)
		}
	}

	generateRoutes(projectType, targetPath, httpRouter)
	generateMicroserviceMainFiles(projectType, targetPath, database, httpRouter)
	generateDockerCompose(projectType, projectName, database)
	replacePlaceholders(targetPath, projectName, database, httpRouter, authMethod)

	color.Magenta("Project %s has been successfully created!\n", projectName)
}

func generateMicroserviceMainFiles(projectType, projectPath, database, httpRouter string) {
	microservices := getMicroservices(projectType)

	for _, service := range microservices {
		servicePath := filepath.Join(projectPath, service)
		mainFilePath := filepath.Join(servicePath, "main.go")

		templatePath := "internal/templates/main.go.tmpl"
		content, err := os.ReadFile(templatePath)
		if err != nil {
			color.Red("Failed to read template: %v", err)
			return
		}

		routerImport, routerInit, serverStart := getRouterInit(httpRouter)
		databaseImport := getDatabaseImport(service, database)

		updatedContent := strings.ReplaceAll(string(content), "{{PROJECT_NAME}}", filepath.Base(projectPath))
		updatedContent = strings.ReplaceAll(updatedContent, "{{SERVICE_NAME}}", service)
		updatedContent = strings.ReplaceAll(updatedContent, "{{DATABASE_IMPORT}}", databaseImport)
		updatedContent = strings.ReplaceAll(updatedContent, "{{ROUTER_IMPORT}}", routerImport)
		updatedContent = strings.ReplaceAll(updatedContent, "{{ROUTER_INIT}}", routerInit)
		updatedContent = strings.ReplaceAll(updatedContent, "{{SERVER_START}}", serverStart)

		err = os.WriteFile(mainFilePath, []byte(updatedContent), 0644)
		if err != nil {
			color.Red("Failed to create main.go for %s: %v", service, err)
			return
		}
	}
}

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

func replacePlaceholders(rootPath, projectName, database, httpRouter, authMethod string) {
	filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		routerImport, routerInit, serverStart := getRouterInit(httpRouter)

		updatedContent := strings.ReplaceAll(string(content), "{{PROJECT_NAME}}", projectName)
		updatedContent = strings.ReplaceAll(updatedContent, "{{DATABASE}}", database)
		updatedContent = strings.ReplaceAll(updatedContent, "{{HTTP_ROUTER}}", httpRouter)
		updatedContent = strings.ReplaceAll(updatedContent, "{{AUTH_METHOD}}", authMethod)
		updatedContent = strings.ReplaceAll(updatedContent, "{{ROUTER_IMPORT}}", routerImport)
		updatedContent = strings.ReplaceAll(updatedContent, "{{ROUTER_INIT}}", routerInit)
		updatedContent = strings.ReplaceAll(updatedContent, "{{SERVER_START}}", serverStart)

		segments := strings.Split(path, string(os.PathSeparator))
		if len(segments) > 1 {
			serviceName := segments[len(segments)-2]
			updatedContent = strings.ReplaceAll(updatedContent, "{{SERVICE_NAME}}", serviceName)
		}

		return ioutil.WriteFile(path, []byte(updatedContent), info.Mode())
	})
}
