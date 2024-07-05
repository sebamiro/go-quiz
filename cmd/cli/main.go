package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sebamiro/go-quiz/database"
	"github.com/sebamiro/go-quiz/dto"
	"github.com/spf13/cobra"
)

const API_URL = "http://127.0.1:3000/"

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
		Short: "Starts specified quiz. Use <available> to check options",
		Args:  cobra.MinimumNArgs(1),
		Run:   start,
	}

	var cmdLeaderboard = &cobra.Command{
		Use:   "leaderboard [uint quiz id]",
		Short: "leaderboard shows the current quiz leaderboard for the specified quis",
		Args:  cobra.MinimumNArgs(1),
		Run:   leaderboard,
	}

	rootCmd.AddCommand(cmdAvailable, cmdStart, cmdLeaderboard)
	rootCmd.Execute()
}

func start(_ *cobra.Command, args []string) {
	questions, err := getQuizQuestions(args[0])
	if err != nil {
		return
	}

	var quizSubmit dto.QuizSubmit
	for _, q := range questions {
		a := getAnswer(q, false)
		fmt.Println("")
		quizSubmit.Aswers = append(quizSubmit.Aswers, a)
	}
	fmt.Print("Introduce your name: ")
	fmt.Scan(&quizSubmit.Name)

	jsonData, err := json.Marshal(quizSubmit)
	if err != nil {
		return
	}

	resp, err := http.Post(API_URL+"quiz/"+args[0], "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var quizEnd dto.QuizEnd
	err = json.NewDecoder(resp.Body).Decode(&quizEnd)
	fmt.Println(quizEnd)
}

func getQuizQuestions(id string) ([]database.QuizQuestion, error) {
	resp, err := http.Get(API_URL + "quiz/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var questions []database.QuizQuestion
	err = json.NewDecoder(resp.Body).Decode(&questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func getAnswer(q database.QuizQuestion, e bool) uint {
	var a uint
	if e {
		fmt.Println("Please introduce a valid answer number")
	}
	fmt.Print(q)
	fmt.Print("Select the number of your answer: ")
	_, err := fmt.Scanln(&a)
	if err != nil && err.Error() != "unexpected newline" {
		if err.Error() == "EOF" {
			os.Exit(1)
		}
		var discard string
		fmt.Scanln(&discard)
	}
	if err != nil || a > 3 {
		return getAnswer(q, true)
	}
	return a
}

func availables(_ *cobra.Command, _ []string) {
	resp, err := http.Get(API_URL + "quiz/")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var quizes []database.Quiz
	err = json.NewDecoder(resp.Body).Decode(&quizes)
	if err != nil {
		return
	}
	fmt.Println("Cuerrent available quizes are:\n ")
	fmt.Println("[ ID ] | TITLE")
	for _, q := range quizes {
		fmt.Printf("[ %d ] %s\n", q.ID, q.Title)
	}
	fmt.Println("\nTo start a quiz use: \n\tquiz-cli start [ quizid ]")
}

func leaderboard(_ *cobra.Command, args []string) {
	resp, err := http.Get(API_URL + "quiz/" + args[0] + "/leaderboard")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var leaderboard []database.QuizResponse
	err = json.NewDecoder(resp.Body).Decode(&leaderboard)
	if err != nil {
		return
	}
	fmt.Println("Leaderboard:")
	for i, q := range leaderboard {
		fmt.Printf("[ %d ] %s : %d points\n", i+1, q.Username, q.Points)
	}
}
