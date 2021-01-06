package repository

import (
	"time-track-api/model"
	"time-track-api/value"
)

type TimeLogRepository interface {
	FetchByDate(date value.Date) (model.TimeLogCollection, error)
}
