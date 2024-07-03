package database

import "errors"

var userId uint = 0

type Database struct {
	Quizes []Quiz

	// Relation QuizId - QuizQuestions
	QuizQuestions map[uint][]QuizQuestion
	// Relation QuizId - QuizResponse
	QuizResponses map[uint][]QuizResponse
}

func NewDatabase() Database {
	return Database{
		Quizes: QUIZES,
		QuizQuestions: map[uint][]QuizQuestion{
			1: GOLANG_QUIZ,
			2: BARCELONA_QUIZ,
		},
		QuizResponses: map[uint][]QuizResponse{
			1: make([]QuizResponse, 0),
			2: make([]QuizResponse, 0),
		},
	}
}

func (d Database) GetQuizes() []Quiz {
	return d.Quizes
}

func (d Database) GetQuizesById(id uint) (*Quiz, error) {
	for _, q := range d.Quizes {
		if q.ID == id {
			return &q, nil
		}
	}
	return nil, errors.New("Quiz not found")
}

func (d Database) GetQuizQuestions(id uint) ([]QuizQuestion, error) {
	q, ok := d.QuizQuestions[id]
	if !ok {
		return nil, errors.New("Questions not found for such quiz")
	}
	return q, nil
}

func (d Database) GetQuizResponses(id uint) ([]QuizResponse, error) {
	q, ok := d.QuizResponses[id]
	if !ok {
		return nil, errors.New("Responses not found for such quiz")
	}
	return q, nil
}

func (d *Database) AddQuizResponse(id uint, qr QuizResponse) error {
	q, ok := d.QuizResponses[id]
	if !ok {
		return errors.New("Responses not found for such quiz")
	}
	q = append(q, qr)
	return nil
}

var (
	QUIZES = []Quiz{
		{ID: 1, Title: "Golang quiz"},
		{ID: 2, Title: "Barcelona quiz"},
	}

	GOLANG_QUIZ = []QuizQuestion{
		{ID: 1, QuizID: 1, Question: "", Answers: []string{}, CorrectAnswer: 0},
		{ID: 2, QuizID: 1, Question: "", Answers: []string{}, CorrectAnswer: 0},
		{ID: 3, QuizID: 1, Question: "", Answers: []string{}, CorrectAnswer: 0},
		{ID: 4, QuizID: 1, Question: "", Answers: []string{}, CorrectAnswer: 0},
		{ID: 5, QuizID: 1, Question: "", Answers: []string{}, CorrectAnswer: 0},
	}

	BARCELONA_QUIZ = []QuizQuestion{
		{ID: 1, QuizID: 2, Question: "", Answers: []string{}, CorrectAnswer: 0},
		{ID: 2, QuizID: 2, Question: "", Answers: []string{}, CorrectAnswer: 0},
		{ID: 3, QuizID: 2, Question: "", Answers: []string{}, CorrectAnswer: 0},
		{ID: 4, QuizID: 2, Question: "", Answers: []string{}, CorrectAnswer: 0},
		{ID: 5, QuizID: 2, Question: "", Answers: []string{}, CorrectAnswer: 0},
	}
)
