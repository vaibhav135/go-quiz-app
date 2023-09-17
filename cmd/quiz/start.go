package quiz

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz/database"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz/gui"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz and check your skillz...",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if timerDuration < 1 {
			return fmt.Errorf("\n Timer duration cannot be less than 1 minute \n")
		}

		if totalQuestions < 5 {
			return fmt.Errorf("\n Total questions can't be less than 5 \n")
		}

		// database.
		quizContent := database.QuizInstance.List(totalQuestions)

		// Since timeDuration is in seconds we are converting it to minutes.
		timerDurationInMinutes := timerDuration * 60
		gui.QuizGUI(timerDurationInMinutes, quizContent)

		fmt.Println(quizContent)
		return nil
	},
}

var timerDuration, totalQuestions int

func init() {
	startCmd.Flags().IntVarP(&timerDuration, "timer", "t", 1, "Timer for the quiz (time is in mins)")
	startCmd.Flags().IntVarP(&totalQuestions, "totalQuestions", "q", 10, "Total no. of questions you want to have in the quiz")

	rootCmd.AddCommand(startCmd)
}
