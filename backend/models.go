package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	ID           int    `json:"id"`
	SubjectID    int    `json:"subject_id"`
	QuestionText string `json:"question_text"`
	AnswerText   string `json:"answer_text,omitempty"`
}

type Answer struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"question_id"`
	AnswerText string `json:"answer_text"`
}

func initDB() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Successfully connected to the database")
}

func getAllSubjects() ([]Subject, error) {
	rows, err := db.Query("SELECT id, name FROM subjects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []Subject
	for rows.Next() {
		var s Subject
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		subjects = append(subjects, s)
	}
	return subjects, nil
}

func getQuestionsBySubjectID(subjectID int) ([]Question, error) {
	query := `
    SELECT q.id, q.subject_id, q.question_text, a.answer_text
    FROM questions q
    LEFT JOIN answers a ON q.id = a.question_id
    WHERE q.subject_id = ?
    `
	rows, err := db.Query(query, subjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var q Question
		var answerText sql.NullString
		if err := rows.Scan(&q.ID, &q.SubjectID, &q.QuestionText, &answerText); err != nil {
			return nil, err
		}
		if answerText.Valid {
			q.AnswerText = answerText.String
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func insertQuestion(subjectID int, questionText string) (int, error) {
	result, err := db.Exec("INSERT INTO questions (subject_id, question_text) VALUES (?, ?)", subjectID, questionText)
	if err != nil {
		return 0, fmt.Errorf("failed to insert question: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %v", err)
	}

	return int(id), nil
}

func getQuestionByID(questionID int) (Question, error) {
	var q Question
	query := `
    SELECT q.id, q.subject_id, q.question_text, a.answer_text
    FROM questions q
    LEFT JOIN answers a ON q.id = a.question_id
    WHERE q.id = ?
    `
	err := db.QueryRow(query, questionID).Scan(&q.ID, &q.SubjectID, &q.QuestionText, &q.AnswerText)
	if err != nil {
		log.Printf("Error querying question by ID %d: %v", questionID, err)
		return Question{}, err
	}
	return q, nil
}

func insertAnswer(a Answer) (int64, error) {
	log.Printf("Inserting answer: %+v", a)
	result, err := db.Exec("INSERT INTO answers (question_id, answer_text) VALUES (?, ?)", a.QuestionID, a.AnswerText)
	if err != nil {
		log.Printf("Error executing insert query: %v", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return 0, err
	}
	log.Printf("Answer inserted with ID: %d", id)
	return id, nil
}

func removeQuestion(questionID int) error {
	_, err := db.Exec("DELETE FROM questions WHERE id = ?", questionID)
	return err
}
