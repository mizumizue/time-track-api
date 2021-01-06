package model

type TimeLogCollection []*TimeLog

type TimeLog struct {
	ID      string  `json:"id"`
	TaskID  string  `json:"taskId"`
	UserID  string  `json:"userId"`
	Hours   float64 `json:"hours"`
	Created string  `json:"createdDate"`
	Updated string  `json:"updatedDate"`
	Tracked string  `json:"trackedDate"`
	Comment string  `json:"comment"`
}
