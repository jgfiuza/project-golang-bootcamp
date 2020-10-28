package qanda

import (
	"context"
	"errors"
)

// QuestionService exposes
type QuestionService interface {
	Create(ctx context.Context, question Question) (Question, error)
	GetAll(ctx context.Context) []Question
	GetAllByUser(ctx context.Context, userID string) []Question
	GetOneByID(ctx context.Context, questionID string) (Question, error)
	Update(ctx context.Context, questionID string, question Question) error
	Delete(ctx context.Context, questionID string) error
}

// AnswerService exposes
type AnswerService interface {
	Create(ctx context.Context, answer Answer) (Answer, error)
	GetAllByQuestion(ctx context.Context, questionID string) []Answer
	GetOneByQuestion(ctx context.Context, questionID string) (Answer, error)
	Delete(ctx context.Context, answerID string) error
}

var (
	// ErrNotFound is thrown when resource with ID doesn't exist.
	ErrNotFound = errors.New("Not found")
)
