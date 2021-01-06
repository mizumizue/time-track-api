package model

type TaskCollection []*Task

// TODO 必要なフィールドの精査
// TODO add json tag
type Task struct {
	ID             string
	AccountID      string
	Title          string
	Status         string
	Importance     string
	CreatedDate    string
	UpdatedDate    string
	Scope          string
	CustomStatusID string
	PermaLink      string
	Priority       string
}
