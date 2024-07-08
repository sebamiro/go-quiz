package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/sebamiro/go-quiz/database"
	"github.com/sebamiro/go-quiz/pkg/dto"
	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, args []string) {
	questions, err := getQuizQuestions(args[0])
	if err != nil {
		cmd.PrintErrln("Error start:", err)
		return
	}

	var quizSubmit dto.QuizSubmit
	for _, q := range questions {
		a := getAnswer(q, false)
		fmt.Print("\n")
		quizSubmit.Answers = append(quizSubmit.Answers, a)
	}
	fmt.Print("Introduce your name: ")
	fmt.Scan(&quizSubmit.Name)

	jsonData, err := json.Marshal(quizSubmit)
	if err != nil {
		cmd.PrintErrln("Error start:", err)
		return
	}

	resp, err := http.Post(API_URL+"quiz/"+args[0], "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		cmd.PrintErrln("Error start:", err)
		return
	}
	defer resp.Body.Close()

	var quizEnd dto.ResponseEnd
	err = json.NewDecoder(resp.Body).Decode(&quizEnd)
	if err != nil {
		cmd.PrintErrln("Error start:", err)
		return
	}
	if quizEnd.Error != "" {
		cmd.PrintErrln("Error start:", quizEnd.Error)
		return
	}
	fmt.Println(quizEnd)
}

func getQuizQuestions(id string) ([]database.QuizQuestion, error) {
	resp, err := http.Get(API_URL + "quiz/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var questions dto.ResopnseQuestions
	err = json.NewDecoder(resp.Body).Decode(&questions)
	if err != nil {
		return nil, err
	}
	if questions.Error != "" {
		return nil, errors.New(questions.Error)
	}
	return questions.Questions, nil
}

func getAnswer(q database.QuizQuestion, e bool) uint {
	var a uint
	if e {
		fmt.Println("\nPlease introduce a valid number as answer")
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
