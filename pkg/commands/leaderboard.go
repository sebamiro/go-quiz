package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sebamiro/go-quiz/pkg/dto"
	"github.com/spf13/cobra"
)

func leaderboard(cmd *cobra.Command, args []string) {
	resp, err := http.Get(API_URL + "quiz/" + args[0] + "/leaderboard")
	if err != nil {
		cmd.PrintErrln("Error leaderboard:", err)
		return
	}
	defer resp.Body.Close()

	var leaderboard dto.ResopnseLeaderboard
	err = json.NewDecoder(resp.Body).Decode(&leaderboard)
	if err != nil {
		cmd.PrintErrln("Error leaderboard:", err)
		return
	}
	if leaderboard.Error != "" {
		cmd.PrintErrln("Error leaderboard:", err)
		return
	}
	fmt.Printf("Leaderboard for %s:\n", leaderboard.Title)
	for i, q := range leaderboard.Leaderboard {
		fmt.Printf("[ %d ] %s : %d points\n", i+1, q.Username, q.Points)
	}
}
