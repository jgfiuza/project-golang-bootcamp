package qanda

// Question is the model for the 'q' part on q&a
type Question struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatorID string `json:"created_by"`
}

// Answer is the model for the 'a' part on q&a
type Answer struct {
	ID         string `json:"id"`
	QuestionID string `json:"question_id"`
	Body       string `json:"body"`
	CreatorID  string `json:"created_by"`
}
