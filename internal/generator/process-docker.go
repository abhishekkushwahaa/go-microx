package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func generateDockerCompose(projectType, projectName, database string) {
	microservices := getMicroservices(projectType)
	composePath := filepath.Join(projectName, "docker-compose.yml")

	var composeContent strings.Builder
	composeContent.WriteString("version: '3.8'\n\nservices:\n")

	for _, service := range microservices {
		composeContent.WriteString(fmt.Sprintf("  %s:\n", service))
		composeContent.WriteString("    build: ./\n")
		composeContent.WriteString(fmt.Sprintf("    working_dir: /app/%s\n", service))
		composeContent.WriteString(fmt.Sprintf("    command: [\"./app\"]\n"))
		composeContent.WriteString("    ports:\n")
		composeContent.WriteString(fmt.Sprintf("      - \"%d:8080\"\n", getPort(service)))
		composeContent.WriteString("    depends_on:\n")
		composeContent.WriteString("      - " + service + "-db\n\n")

		if database != "None" {
			composeContent.WriteString(fmt.Sprintf("  %s-db:\n", service))
			composeContent.WriteString("    image: ")

			switch database {
			case "PostgreSQL":
				composeContent.WriteString("postgres:15-alpine\n")
				composeContent.WriteString("    environment:\n")
				composeContent.WriteString("      POSTGRES_USER: user\n      POSTGRES_PASSWORD: password\n      POSTGRES_DB: " + service + "\n")
				composeContent.WriteString("    ports:\n")
				composeContent.WriteString(fmt.Sprintf("      - \"%d:5432\"\n", getDBPort(service)))
			case "MongoDB":
				composeContent.WriteString("mongo:6.0\n")
				composeContent.WriteString("    ports:\n")
				composeContent.WriteString(fmt.Sprintf("      - \"%d:27017\"\n", getDBPort(service)))
			case "MySQL":
				composeContent.WriteString("mysql:8.0\n")
				composeContent.WriteString("    environment:\n")
				composeContent.WriteString("      MYSQL_ROOT_PASSWORD: password\n      MYSQL_DATABASE: " + service + "\n")
				composeContent.WriteString("    ports:\n")
				composeContent.WriteString(fmt.Sprintf("      - \"%d:3306\"\n", getDBPort(service)))
			}
			composeContent.WriteString("\n")
		}
	}

	err := os.WriteFile(composePath, []byte(composeContent.String()), 0644)
	if err != nil {
		color.Red("Failed to create docker-compose.yml: %v", err)
		return
	}
}

func getPort(service string) int {
	portMap := map[string]int{
		"user-service":           8081,
		"product-service":        8082,
		"order-service":          8083,
		"payment-service":        8084,
		"cart-service":           8085,
		"video-service":          8086,
		"subscription-service":   8087,
		"recommendation-service": 8088,
		"analytics-service":      8089,
		"restaurant-service":     8090,
		"delivery-service":       8091,
		"api-service":            8092,
	}
	return portMap[service]
}

func getDBPort(service string) int {
	return getPort(service) + 1000
}
