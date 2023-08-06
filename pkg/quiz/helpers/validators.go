package helpers

import (
	fp "path/filepath"
	"strings"

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

func ValidateFileExtensions(filePath string) bool {
  var fileExtension = fp.Ext(filePath)

  var fileType = strings.Split(fileExtension, ".")[1]
  
  return ValidateFiletype(fileType)
}
