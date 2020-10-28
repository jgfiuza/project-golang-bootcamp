package qanda

// Question is the model for the 'q' part on q&a
type Question struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatorID string `json:"question_by"`
}

// Answer is the model for the 'a' part on q&a
type Answer struct {
	ID         string `json:"id"`
	Body       string `json:"body"`
	QuestionID string `json:"question_id"`
	CreatorID  string `json:"answer_by"`
	Version    int    `json:"version"`
}
