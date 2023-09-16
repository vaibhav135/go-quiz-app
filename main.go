package main

import (
	"github.com/vaibhav135/go-quiz-app/cmd/quiz"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz/database"
)

func main() {
	conn := database.IntializeQuiz()

	quiz.Execute()

	defer conn.CloseConn()
}
