package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sebamiro/go-quiz/database"
	"github.com/spf13/cobra"
)

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
