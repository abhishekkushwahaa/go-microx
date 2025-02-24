package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new microservices project",
	Run: func(cmd *cobra.Command, args []string) {
		projectType := selectProjectType()
		projectName := getProjectName()
		database := selectDatabase()

		color.Cyan("\n🚀 Creating project: %s", projectName)
		color.Green("📦 Type: %s", projectType)
		color.Yellow("🫙  Database: %s\n", database)

		createProjectStructure(projectType, projectName, database)
		color.Magenta("✅ Project %s has been successfully created!\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

// Prompt user to select a project type
func selectProjectType() string {
	templates := []string{"E-commerce", "Video Streaming", "Food Delivery", "Custom"}

	prompt := promptui.Select{
		Label:  "💡 " + color.HiCyanString("Select a project template:"),
		Items:  templates,
		Stdout: os.Stderr,
	}

	_, result, err := prompt.Run()
	if err != nil {
		color.Red("❌ Selection failed: %v", err)
		os.Exit(1)
	}

	return result
}

// Prompt user to enter a project name
func getProjectName() string {
	prompt := promptui.Prompt{
		Label:  color.HiBlueString("Enter project name:"),
		Stdout: os.Stderr,
	}

	result, err := prompt.Run()
	if err != nil {
		color.Red("❌ Project name input failed: %v", err)
		os.Exit(1)
	}

	return result
}

// Prompt user to select a database
func selectDatabase() string {
	databases := []string{"PostgreSQL", "MongoDB", "MySQL", "SQLite", "None"}

	prompt := promptui.Select{
		Label:  color.HiYellowString("Select a database:"),
		Items:  databases,
		Stdout: os.Stderr,
	}

	_, result, err := prompt.Run()
	if err != nil {
		color.Red("❌ Database selection failed: %v", err)
		os.Exit(1)
	}

	return result
}

// Create project structure based on selections
func createProjectStructure(projectType, projectName, database string) {
	basePath := filepath.Join(".", projectName)
	os.MkdirAll(basePath, os.ModePerm)

	services := getServicesForProject(projectType)
	for _, service := range services {
		servicePath := filepath.Join(basePath, service)
		os.MkdirAll(servicePath, os.ModePerm)

		mainFile := filepath.Join(servicePath, "main.go")
		err := os.WriteFile(mainFile, []byte(fmt.Sprintf("package main\n\nfunc main() {\n\tprintln(\"%s service running...\")\n}", service)), 0644)
		if err != nil {
			color.Red("❌ Error creating service file: %v", err)
		}
	}

	// Create database configuration file
	if database != "None" {
		configPath := filepath.Join(basePath, "configs")
		os.MkdirAll(configPath, os.ModePerm)

		dbConfigFile := filepath.Join(configPath, "db_config.yaml")
		dbConfigContent := fmt.Sprintf("database: %s\nhost: localhost\nport: 5432\nuser: root\npassword: secret", database)
		os.WriteFile(dbConfigFile, []byte(dbConfigContent), 0644)
	}
}

// Get predefined services based on project type
func getServicesForProject(projectType string) []string {
	switch projectType {
	case "E-commerce":
		return []string{"user-service", "product-service", "order-service", "cart-service", "payment-service"}
	case "Video Streaming":
		return []string{"video-service", "user-service", "subscription-service"}
	case "Food Delivery":
		return []string{"restaurant-service", "order-service", "delivery-service"}
	default:
		return []string{"custom-service1", "custom-service2"}
	}
}
