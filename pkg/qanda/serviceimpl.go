package qanda

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

// *** QUESTION SERVICE ***
// TODO: substitute for a real life RDBMS
// Makes an inmemory question service that provides the interfaces methods
type inMemoryQuestionService struct {
	sync.RWMutex
	m map[string]Question
}

// method implementation
func (q *inMemoryQuestionService) Create(ctx context.Context, question Question) (Question, error) {
	q.Lock()
	defer q.Unlock()

	question.ID = uuid.New().String()
	q.m[question.ID] = question

	return question, nil
}

func (q *inMemoryQuestionService) GetAll(ctx context.Context) []Question {
	q.Lock()
	defer q.Unlock()

	var questionList []Question

	for _, question := range q.m {
		questionList = append(questionList, question)
	}
	return questionList
}

func (q *inMemoryQuestionService) GetAllByUser(ctx context.Context, userID string) []Question {
	q.Lock()
	defer q.Unlock()

	var questionList []Question

	for _, question := range q.m {
		if question.CreatorID == userID {
			questionList = append(questionList, question)
		}
	}

	return questionList
}

func (q *inMemoryQuestionService) GetOneByID(ctx context.Context, questionID string) (Question, error) {
	q.Lock()
	defer q.Unlock()

	if question, ok := q.m[questionID]; ok {
		return question, nil
	}

	return Question{}, ErrNotFound
}

func (q *inMemoryQuestionService) Update(ctx context.Context, questionID string, question Question) error {
	q.Lock()
	defer q.Unlock()

	if _, ok := q.m[questionID]; !ok {
		return ErrNotFound
	}

	q.m[questionID] = question

	return nil
}

func (q *inMemoryQuestionService) Delete(ctx context.Context, questionID string) error {
	q.Lock()
	defer q.Unlock()

	if _, ok := q.m[questionID]; !ok {
		return ErrNotFound
	}

	delete(q.m, questionID)

	return nil
}

// *** ANSWER SERVICE ***
// TODO: substitute for a real life RDBMS
// TODO: implement version control
// Makes an inmemory question service that provides the interfaces methods
type inMemoryAnswerService struct {
	sync.RWMutex
	m map[string]Answer
}

func (a *inMemoryAnswerService) Create(ctx context.Context, answer Answer) (Answer, error) {
	a.Lock()
	defer a.Unlock()

	answer.ID = uuid.New().String()

	a.m[answer.ID] = answer
	return answer, nil
}

func (a *inMemoryAnswerService) GetAllByQuestion(ctx context.Context, questionID string) []Answer {
	a.Lock()
	defer a.Unlock()

	var answerList []Answer

	for _, answer := range a.m {
		if answer.QuestionID == questionID {
			answerList = append(answerList, answer)
		}
	}

	return answerList
}

func (a *inMemoryAnswerService) GetOneByQuestion(ctx context.Context, questionID string) (Answer, error) {
	a.Lock()
	defer a.Unlock()

	// TODO: implement an algorithm that returns the last version only
	for _, ans := range a.m {
		if ans.QuestionID == questionID {
			return ans, nil
		}
	}

	return Answer{}, nil
}

func (a *inMemoryAnswerService) Delete(ctx context.Context, answerID string) error {
	a.Lock()
	defer a.Unlock()

	if _, ok := a.m[answerID]; !ok {
		return ErrNotFound
	}

	delete(a.m, answerID)

	return nil
}
