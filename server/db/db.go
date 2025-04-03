package db

import (
	"database/sql"
	"sync"
	"duabi/models"
)

type DBManager struct {
	DB *sql.DB
	WG *sync.WaitGroup
	MU *sync.RWMutex
}

func NewDBManager(driver string, connStr string) (*DBManager, error) {
	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	return &DBManager{DB: db, WG: wg, MU: mu}, nil
}

func (d *DBManager) Close() {
	d.DB.Close()
	d.DB = nil
}

func (d *DBManager) GetQuestions(categoryId int) ([]models.Questions, error) {
	var questions []models.Questions
	query := "SELECT id, question FROM questions WHERE category_id = $1"
	rows, err := d.DB.Query(query, categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question models.Questions
		if err := rows.Scan(&question.ID, &question.Question); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func (d *DBManager) GetAnswer(questionId int) (string, error) {
	var question string
	
	query := "SELECT answer FROM answers WHERE question_id = $1"
	row := d.DB.QueryRow(query, questionId)
	if err := row.Scan(&question); err != nil {
		return "", err
	}

	return question, nil
}