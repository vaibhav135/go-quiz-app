package helpers 

import (
	"fmt"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz"
)

func ParseFile(fileType string) bool {
  switch fileType{
    case quiz.Json:
      fmt.Println("parsing a json")
    default:
      fmt.Println("parsing a csv")
  }
  
  return false 
}

