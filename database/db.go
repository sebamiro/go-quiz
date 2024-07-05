package database

import (
	"errors"
	"sort"
)

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
	d.QuizResponses[id] = append(q, qr)
	return nil
}

func (d *Database) GetQuizResponsesOrderdByPoints(id uint) ([]QuizResponse, error) {
	q, ok := d.QuizResponses[id]
	if !ok {
		return nil, errors.New("Responses not found for such quiz")
	}
	sort.Sort(sort.Reverse(ResponsesList(q)))
	return q, nil
}

var (
	QUIZES = []Quiz{
		{ID: 1, Title: "Golang quiz"},
		{ID: 2, Title: "Barcelona quiz"},
	}

	GOLANG_QUIZ = []QuizQuestion{
		{
			ID:       1,
			QuizID:   1,
			Question: "What keyword is used to define a package in Go?",
			Answers: []string{
				"namespace",
				"package",
				"module",
				"library",
			},
			CorrectAnswer: 1,
		},
		{
			ID:       2,
			QuizID:   1,
			Question: "How do you declare a variable in Go?",
			Answers: []string{
				"var x int",
				"int x",
				"declare x int",
				"define x int",
			},
			CorrectAnswer: 0,
		},
		{
			ID:       3,
			QuizID:   1,
			Question: "What is the correct way to create a slice in Go?",
			Answers: []string{
				"var s []int",
				"var s int[]",
				"var int[] s",
				"s := []int",
			},
			CorrectAnswer: 0,
		},
		{
			ID:       4,
			QuizID:   1,
			Question: "Which of the following is the correct way to handle errors in Go?",
			Answers: []string{
				"try { ... } catch { ... }",
				"expect { ... }",
				"if err != nil { ... }",
				"error { ... }",
			},
			CorrectAnswer: 2,
		},
		{
			ID:       5,
			QuizID:   1,
			Question: "What is the name of the tool used for formatting Go code?",
			Answers: []string{
				"goformat",
				"go-fmt",
				"formatgo",
				"gofmt",
			},
			CorrectAnswer: 3,
		},
	}

	BARCELONA_QUIZ = []QuizQuestion{
		{
			ID:       1,
			QuizID:   2,
			Question: "What is the name of the famous basilica in Barcelona designed by Antoni Gaudí?",
			Answers: []string{
				"La Sagrada Familia",
				"La Catedral",
				"La Almudena",
				"La Giralda",
			},
			CorrectAnswer: 0,
		},
		{
			ID:       2,
			QuizID:   2,
			Question: "Which of the following is a famous street in Barcelona known for its shops and cafes?",
			Answers: []string{
				"Champs-Élysées",
				"Oxford Street",
				"Fifth Avenue",
				"Las Ramblas",
			},
			CorrectAnswer: 3,
		},
		{
			ID:       3,
			QuizID:   2,
			Question: "Which park in Barcelona is known for its mosaics and architecture by Gaudí?",
			Answers: []string{
				"Retiro Park",
				"Hyde Park",
				"Park Güell",
				"Central Park",
			},
			CorrectAnswer: 2,
		},
		{
			ID:       4,
			QuizID:   2,
			Question: "What is the name of Barcelona's main football stadium?",
			Answers: []string{
				"Santiago Bernabéu Stadium",
				"Camp Nou",
				"San Siuro",
				"Allianz Arena",
			},
			CorrectAnswer: 1,
		},
		{
			ID:       5,
			QuizID:   1,
			Question: "Which historic neighborhood in Barcelona is known for its Gothic architecture and narrow streets?",
			Answers: []string{
				"El Raval",
				"Gràcia",
				"El Born",
				"Barri Gothic",
			},
			CorrectAnswer: 3,
		},
	}
)

type ResponsesList []QuizResponse

func (p ResponsesList) Len() int           { return len(p) }
func (p ResponsesList) Less(i, j int) bool { return p[i].Points < p[j].Points }
func (p ResponsesList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
