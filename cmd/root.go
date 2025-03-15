package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-microx",
	Short: "go-microx - Microservices Generator",
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan).SprintFunc()
		fmt.Println(cyan("MicroX - Generate microservices projects quickly!"))
	},
}

func Execute() error {
	return rootCmd.Execute()
}
