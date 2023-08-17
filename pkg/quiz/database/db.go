package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz"
)


type Quiz struct {
  DB *sql.DB
}

var QuizInstance *Quiz
var dbFile = "../../../quiz.db"
var lock = &sync.Mutex{}

func IntializeQuiz() (*Quiz, error){

  lock.Lock()
  defer lock.Unlock()

  if QuizInstance == nil {
    db, err := sql.Open("sqlite3", dbFile)

    if err != nil {
      return nil, err 
    }

    if _, err := db.Exec(dbSeedQuery); err != nil {
      return nil, err
    }

    QuizInstance = &Quiz{
      DB: db,
    }

  }else{
    log.Printf("Instance already created...")
  }

  return QuizInstance, nil
}

func (quiz *Quiz) List(numberOfQuiz int) {

}


func (quiz *Quiz) BulkInsert(data []quiz.QuizContent) {

  for _, question := range data{
    err := quiz.Insert(question) 
    if err != nil {
      log.Fatal(err)
    }
  }

  log.Println("\nData added successfully...")
}

func (quiz *Quiz) Insert(data quiz.QuizContent) error {
  _, err := quiz.DB.Exec(quizInsertQuery, data.Question, data.Answer)

  if err != nil {
    return err
  }

  return nil
}

func (quiz *Quiz) CloseConn() {
  QuizInstance.DB.Close()
  QuizInstance = nil
}
