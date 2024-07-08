package commands

import "github.com/spf13/cobra"

const API_URL = "http://127.0.0.1:3000/"

func BuildCommands(rootCmd *cobra.Command) {
	var cmdAvailable = &cobra.Command{
		Use:   "available",
		Short: "Check quizes available",
		Args:  cobra.MinimumNArgs(0),
		Run:   availables,
	}

	var cmdStart = &cobra.Command{
		Use:   "start [uint quiz id]",
		Short: "Starts specified quiz. Use <available> to check options",
		Args:  cobra.MinimumNArgs(1),
		Run:   start,
	}

	var cmdLeaderboard = &cobra.Command{
		Use:   "leaderboard [uint quiz id]",
		Short: "leaderboard shows the current quizers leaderboard for the specified quiz",
		Args:  cobra.MinimumNArgs(1),
		Run:   leaderboard,
	}

	rootCmd.AddCommand(cmdAvailable, cmdStart, cmdLeaderboard)
}
