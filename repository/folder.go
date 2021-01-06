package repository

import "time-track-api/model"

type FolderRepository interface {
	FetchByIDs(IdList []string) (model.FolderCollection, error)
}
