package cmd

import (
	"os"

	"github.com/abhishekkushwahaa/go-microx/cmd/generator"
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

		generator.GenerateProject(projectType, projectName, database)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

// Prompt user to select a project type
func selectProjectType() string {
	templates := []string{"E-commerce", "Video-Streaming", "Food-Delivery", "Custom"}

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
