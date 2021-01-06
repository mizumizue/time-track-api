package wrikeapi

import (
	"time-track-api/infrastructure/http"
	"time-track-api/model"
	"time-track-api/repository"
)

type FolderAPI struct {
	client http.HttpClient
}

func NewFolderAPI(client http.HttpClient) repository.FolderRepository {
	return &FolderAPI{client: client}
}

func (f *FolderAPI) basePath() string {
	return "https://www.wrike.com/api/v4/folders"
}

func (f *FolderAPI) FetchByIDs(IdList []string) (model.FolderCollection, error) {
	panic("implement me")
}
