package quiz

import (
	"github.com/spf13/cobra"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz/gui"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz and check your skillz...",
	Args: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		gui.Timer(3)

		return nil
	},
}

var timer int

func init() {
	startCmd.Flags().IntVarP(&timer, "timer", "t", 1, "Timer for the quiz (time is in mins)")
	startCmd.Flags().IntVarP(&timer, "totalQuestions", "q", 10, "Total no. of questions you want to have in the quiz")

	rootCmd.AddCommand(startCmd)
}
