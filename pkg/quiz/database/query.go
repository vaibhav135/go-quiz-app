package database

var dbSeedQuery = `CREATE TABLE IF NOT EXISTS quiz (
                    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
                    question TEXT UNIQUE NOT NULL, 
                    answer VARCHAR(100) NOT NULL
                  )`

var quizInsertQuery = `INSERT INTO quiz(question, answer) VALUES(?, ?)`

var quizListQuery = `SELECT question, answer FROM quiz ORDER BY RANDOM() LIMIT ?`
