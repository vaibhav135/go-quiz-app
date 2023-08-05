package quiz

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a quiz using a json or csv file",
  Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) { 
    fmt.Println(args)
    fmt.Println(author, fileType)
  },
}

var author, fileType string

func init() {
	addCmd.Flags().StringVarP(&author, "author", "a", "", "Add name of author who is adding the quiz")
	addCmd.Flags().StringVarP(&fileType,"type", "t", "json", "Add the type of file you want to add -> json | csv")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
  rootCmd.AddCommand(addCmd)
}


