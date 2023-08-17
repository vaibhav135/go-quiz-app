package helpers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/vaibhav135/go-quiz-app/pkg/quiz"
)

func ParseFile(fileType string, filePath string) ([]quiz.QuizContent, error) {
  content, err := os.ReadFile(filePath)
  var quizContent *[]quiz.QuizContent

  if err != nil {
    return *quizContent, err 
  }
  fmt.Println(content)

  switch fileType{
    case quiz.Json:
    err := json.Unmarshal(content, &quizContent)

    if err != nil {
      return  *quizContent, fmt.Errorf("Couldn't parse json. \n%s\n", err)
    }

    default:
      fmt.Println("parsing a csv")
  }
  return *quizContent, nil
}

