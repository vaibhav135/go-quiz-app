package quiz

import (
	"github.com/spf13/cobra"
)


var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz and check your skillz...",
	Run: func(cmd *cobra.Command, args []string) { },
}


func init() {
  rootCmd.AddCommand(startCmd)
}


