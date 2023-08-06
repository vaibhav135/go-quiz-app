package quiz

import (
	"fmt"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz/utility"
)

func parseFile(fileType string) bool {
  switch fileType{
    case Json:
      fmt.Println("parsing a json")
    default:
      fmt.Println("parsing a csv")
  }
  
  return false 
}

func validateFiletype(fileType string) bool {
  var fileTypes = []string {Json, Csv} 

  if utility.Compare(fileType, fileTypes) {
    return true
  }

  return false 
}
