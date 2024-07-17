package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getSubjects(c *gin.Context) {
	subjects, err := getAllSubjects()
	if err != nil {
		log.Printf("Error getting subjects: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subjects)
}

func getQuestionsBySubject(c *gin.Context) {
	subjectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid subject ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	questions, err := getQuestionsBySubjectID(subjectID)
	if err != nil {
		log.Printf("Error getting questions for subject %d: %v", subjectID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Retrieved %d questions for subject %d", len(questions), subjectID)
	c.JSON(http.StatusOK, questions)
}

func createQuestion(c *gin.Context) {
	var newQuestion struct {
		SubjectID int    `json:"subject"`
		Question  string `json:"question"`
	}

	if err := c.BindJSON(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// バリデーション
	if newQuestion.SubjectID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subject ID is required"})
		return
	}
	if newQuestion.Question == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Question text is required"})
		return
	}

	// データベースに質問を保存
	questionID, err := insertQuestion(newQuestion.SubjectID, newQuestion.Question)
	if err != nil {
		log.Printf("Failed to insert question: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": questionID})
}

func getQuestion(c *gin.Context) {
	questionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	question, err := getQuestionWithAnswers(questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, question)
}

func getQuestionWithAnswers(questionID int) (map[string]interface{}, error) {
	query := `
    SELECT q.id, q.subject_id, q.question_text, a.answer_text
    FROM questions q
    LEFT JOIN answers a ON q.id = a.question_id
    WHERE q.id = ?
    `
	rows, err := db.Query(query, questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var question map[string]interface{}
	var answers []string

	for rows.Next() {
		var q Question
		var answerText sql.NullString
		if err := rows.Scan(&q.ID, &q.SubjectID, &q.QuestionText, &answerText); err != nil {
			return nil, err
		}
		if question == nil {
			question = map[string]interface{}{
				"id":            q.ID,
				"subject_id":    q.SubjectID,
				"question_text": q.QuestionText,
			}
		}
		if answerText.Valid {
			answers = append(answers, answerText.String)
		}
	}

	if question != nil {
		question["answers"] = answers
	}

	return question, nil
}

func createAnswer(c *gin.Context) {
	questionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid question ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	var answer Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	answer.QuestionID = questionID

	log.Printf("Attempting to insert answer: %+v", answer)
	id, err := insertAnswer(answer)
	if err != nil {
		log.Printf("Error inserting answer: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Answer inserted successfully with ID: %d", id)
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func deleteQuestion(c *gin.Context) {
	questionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	err = removeQuestion(questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
