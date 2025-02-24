package main

import (
	"fmt"
	"os"

	"github.com/abhishekkushwahaa/go-microx/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
