package helpers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/vaibhav135/go-quiz-app/pkg/quiz"
)

func ParseFile(fileType string, filePath string) ([]quiz.QuizContent, error) {
	content, err := os.ReadFile(filePath)
	var quizContent []quiz.QuizContent

	if err != nil {
		return quizContent, err
	}

	switch fileType {
	case quiz.Json:
		err := json.Unmarshal(content, &quizContent)

		if err != nil {
			return quizContent, fmt.Errorf("Couldn't parse json. \n%s\n", err)
		}

	default:
		csvReader := csv.NewReader(bytes.NewBuffer(content))
		csvData, err := csvReader.ReadAll()

		if err != nil {
			return nil, err
		}

		for _, data := range csvData {
			var quizData quiz.QuizContent

			quizData.Question = data[0]
			quizData.Answer = data[1]

			quizContent = append(quizContent, quizData)
		}

	}
	return quizContent, nil
}
