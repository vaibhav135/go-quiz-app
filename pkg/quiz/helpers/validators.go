package helpers

import (
	"github.com/vaibhav135/go-quiz-app/pkg/quiz"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz/utility"
)

func ValidateFiletype(fileType string) bool {
  var fileTypes = []string {quiz.Json, quiz.Csv} 

  if utility.Compare(fileType, fileTypes) {
    return true
  }

  return false 
}

