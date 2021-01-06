package wrikeapi

import (
	"time-track-api/infrastructure/http"
	"time-track-api/model"
	"time-track-api/repository"
)

type TaskAPI struct {
	client http.HttpClient
}

func NewTaskAPI(client http.HttpClient) repository.TaskRepository {
	return &TaskAPI{client: client}
}

func (t *TaskAPI) basePath() string {
	return "https://www.wrike.com/api/v4/tasks"
}

func (t *TaskAPI) FetchByIDs(IdList []string) (model.TaskCollection, error) {
	panic("implement me")
}
