package commands

import "github.com/spf13/cobra"

func BuildCommands(rootCmd *cobra.Command) {
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
}
