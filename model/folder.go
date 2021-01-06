package model

import "strings"

type FolderCollection []*Folder

// TODO 必要なフィールドの精査
// TODO add json tag
type Folder struct {
	ID          string
	Title       string
	CreatedDate string
	UpdatedDate string
}

func (f *Folder) OrderNumber() string {
	return strings.Split(f.Title, ":")[0]
}
