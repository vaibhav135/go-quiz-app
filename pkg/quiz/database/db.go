package database

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	TQuiz "github.com/vaibhav135/go-quiz-app/pkg/quiz"
)

type Quiz struct {
	DB *sql.DB
}

var QuizInstance *Quiz
var dbFile = "quiz.db"
var lock = &sync.Mutex{}

func IntializeQuiz() *Quiz {

	lock.Lock()
	defer lock.Unlock()

	if QuizInstance == nil {
		db, err := create()

		if err != nil {
			log.Panic(err)
		}

		QuizInstance = &Quiz{
			DB: db,
		}

	} else {
		log.Printf("Instance already created...")
	}

	return QuizInstance
}

func createFile() {
	f, err := os.Create(dbFile)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}

func create() (*sql.DB, error) {
	_, err := os.Stat(dbFile)

	if errors.Is(err, os.ErrNotExist) {
		createFile()
	}

	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(dbSeedQuery); err != nil {
		return nil, err
	}

	return db, err

}

func (quiz *Quiz) List(numberOfQuiz int) []TQuiz.QuizContent {
	rows, err := quiz.DB.Query(quizListQuery, numberOfQuiz)

	if err != nil {
		log.Panic(err)
	}

	var quizContent = []TQuiz.QuizContent{}

	for rows.Next() {
		var (
			Question string
			Answer   string
		)
		if err := rows.Scan(&Question, &Answer); err != nil {
			log.Fatal(err)
		}
		quizContent = append(quizContent, TQuiz.QuizContent{Question: Question, Answer: Answer})
	}

	return quizContent
}

func (quiz *Quiz) BulkInsert(quizData []TQuiz.QuizContent) {
	stmt, err := quiz.DB.Prepare(quizInsertQuery)

	if err != nil {
		log.Panic(err)
	}

	defer stmt.Close()

	for _, data := range quizData {
		_, err := stmt.Exec(data.Question, data.Answer)
		if err != nil {
			log.Panic(err)
		}
	}

	log.Println("\nData added successfully...")
}

func (quiz *Quiz) CloseConn() {
	QuizInstance.DB.Close()
	QuizInstance = nil
}
