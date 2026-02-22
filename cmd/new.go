package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/abhishekkushwahaa/go-microx/internal/generator"
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
		httpRouter := selectHTTPRouter()
		authMethod := selectAuthMethod()

		generator.GenerateProject(projectType, projectName, database, httpRouter, authMethod)
	},
}

const version = "v1.0.0"

var vCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of go-microx",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(vCmd)
}

func selectProjectType() string {
	templates := []string{"E-commerce", "Video-Streaming", "Food-Delivery", "Custom"}

	prompt := promptui.Select{
		Label:  color.HiCyanString("Select a project template:"),
		Items:  templates,
		Stdout: os.Stderr,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(color.RedString("Selection failed: %v", err))
	}

	if result == "Custom" {
		color.Yellow("You selected 'Custom'. Please specify your custom project template.")
		return getCustomProjectTemplate()
	}

	return result
}

func getCustomProjectTemplate() string {
	prompt := promptui.Prompt{
		Label:  color.HiBlueString("Enter custom project template name:"),
		Stdout: os.Stderr,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal(color.RedString("Custom template input failed: %v", err))
	}

	return result
}

func getProjectName() string {
	prompt := promptui.Prompt{
		Label:  color.HiBlueString("Enter project name:"),
		Stdout: os.Stderr,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal(color.RedString("Project name input failed: %v", err))
	}

	return result
}

func selectDatabase() string {
	databases := []string{"PostgreSQL", "MongoDB", "MySQL", "SQLite", "None"}

	prompt := promptui.Select{
		Label:  color.HiYellowString("Select a database:"),
		Items:  databases,
		Stdout: os.Stderr,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(color.RedString("Database selection failed: %v", err))
	}

	if result == "None" {
		color.Yellow("You selected 'None'. No database will be configured.")
	}

	return result
}

func selectHTTPRouter() string {
	routers := []string{"Gin", "Fiber", "Chi", "Echo", "Mux", "Custom"}

	prompt := promptui.Select{
		Label:  color.HiMagentaString("Select an HTTP Router:"),
		Items:  routers,
		Stdout: os.Stderr,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(color.RedString("Router selection failed: %v", err))
	}

	if result == "Custom" {
		color.Yellow("You selected 'Custom'. Please specify your custom HTTP Router.")
		return getCustomHTTPRouter()
	}

	return result
}

func getCustomHTTPRouter() string {
	prompt := promptui.Prompt{
		Label:  color.HiMagentaString("Enter custom HTTP router name:"),
		Stdout: os.Stderr,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal(color.RedString("Custom router input failed: %v", err))
	}

	return result
}

func selectAuthMethod() string {
	authMethods := []string{"None", "JWT", "OAuth", "API Key"}

	prompt := promptui.Select{
		Label:  color.HiGreenString("Select an authentication method:"),
		Items:  authMethods,
		Stdout: os.Stderr,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(color.RedString("Authentication selection failed: %v", err))
	}

	if result == "None" {
		color.Yellow("You selected 'None'. No authentication will be configured.")
	}

	return result
}
