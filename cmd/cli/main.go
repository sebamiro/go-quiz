package main

import (
	"github.com/sebamiro/go-quiz/pkg/commands"
	"github.com/spf13/cobra"
)

const APP_NAME = "quiz-cli"

func main() {
	var rootCmd = &cobra.Command{Use: APP_NAME}

	commands.BuildCommands(rootCmd)
	rootCmd.Execute()
}
