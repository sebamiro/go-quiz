package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sebamiro/go-quiz/database"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "quiz-cli"}

	var cmdAvailable = &cobra.Command{
		Use:   "available",
		Short: "Check quizes available",
		Long: `available is for printing to the screen the current
available quizes with their ids.`,
		Args: cobra.MinimumNArgs(0),
		Run:  availables,
	}

	var cmdStart = &cobra.Command{
		Use:   "start [uint quiz id]",
		Short: "Starts especified quiz. Use <available> to check options",
		Long: `available is for printing to the screen the current
available quizes with their ids.`,
		Args: cobra.MinimumNArgs(1),
		Run:  start,
	}

	rootCmd.AddCommand(cmdAvailable)
	rootCmd.AddCommand(cmdStart)
	rootCmd.Execute()
}

func start(_ *cobra.Command, args []string) {
	resp, err := http.Get("http://localhost:3000/quiz/" + args[0])
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var body []database.QuizQuestion
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return
	}
	fmt.Println(body)
}

func availables(_ *cobra.Command, _ []string) {
	resp, err := http.Get("http://localhost:3000/quiz/")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var body []database.Quiz
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return
	}
	fmt.Println(body)
}
