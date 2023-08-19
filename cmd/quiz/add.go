package quiz

import (
	"errors"
	"fmt"
	fp "path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz/database"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz/helpers"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a quiz using a json or csv file",
  Args: func(cmd *cobra.Command, args []string) error {
    err := cobra.MaximumNArgs(1)(cmd, args)
    if err != nil {
      return err
    }
  
    filePath := args[0]
    var fileExtension = fp.Ext(filePath)
    var fileExtensionType = strings.Split(fileExtension, ".")[1]

    if fileExtensionType != fileType {
      return fmt.Errorf("\n File extension type does not match type flag value.")
    }
    

    if !helpers.ValidateFiletype(fileExtensionType)  {
      return errors.New("\n File extension type can either be .csv or .json")
    }


    return nil
  },
	RunE: func(cmd *cobra.Command, args []string) error{ 

    if !helpers.ValidateFiletype(fileType) {
      return fmt.Errorf("\nFiletype can only be csv or json\n")
    }

    filePath := args[0]

    quizContent, err := helpers.ParseFile(fileType, filePath)
  
    if err != nil {
      return err
    }

    database.QuizInstance.BulkInsert(quizContent) 

    return nil
  },
}

var author, fileType string

func init() {
	addCmd.Flags().StringVarP(&author, "author", "a", "", "Add name of author who is adding the quiz")
	addCmd.Flags().StringVarP(&fileType,"type", "t", "json", "Add the filetype either json | csv")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
  rootCmd.AddCommand(addCmd)
}


