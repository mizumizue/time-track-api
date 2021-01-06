package wrikeapi

import (
	"encoding/json"
	"time-track-api/infrastructure/http"
	"time-track-api/model"
	"time-track-api/value"
)

type TimeLogAPI struct {
	client http.HttpClient
}

func NewTimeLogAPI(client http.HttpClient) *TimeLogAPI {
	return &TimeLogAPI{client: client}
}

func (t *TimeLogAPI) basePath() string {
	return "https://www.wrike.com/api/v4/timelogs?me&trackedDate={\"equal\":\"2021-01-06\"}"
}

func (t *TimeLogAPI) FetchByDate(date value.Date) (collection model.TimeLogCollection, err error) {
	err = func(date value.Date) error {
		req, err := http.NewRequestWithHeader("GET", t.basePath(), nil)
		if err != nil {
			return err
		}
		body, err := t.client.Do(req)
		if err != nil {
			return err
		}
		var data struct {
			Data model.TimeLogCollection `json:"data"`
		}
		if err := json.Unmarshal(body, &data); err != nil {
			return err
		}
		collection = data.Data
		return nil
	}(date)
	return
}
