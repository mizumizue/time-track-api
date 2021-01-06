package repository

import (
	"time-track-api/model"
)

type TaskRepository interface {
	FetchByIDs(IdList []string) (model.TaskCollection, error)
}
