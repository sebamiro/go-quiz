package main

import (
	"os"

	"github.com/sebamiro/go-quiz/pkg/commands"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: os.Args[0]}

	commands.BuildCommands(rootCmd)
	rootCmd.Execute()
}
