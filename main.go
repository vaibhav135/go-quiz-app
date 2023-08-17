/*
Copyright © 2023  <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/vaibhav135/go-quiz-app/cmd/quiz"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz/database"
)

func main() {
  conn, err := database.IntializeQuiz()
  
  if err != nil {
    fmt.Errorf("Sorry couldn't connect to the DB \n %s", err)
  }

	quiz.Execute()
  
  defer conn.CloseConn()
}
