package main

import (
	"log"
	"net/http"
	"time"
	ch "time-track-api/infrastructure/http"
	"time-track-api/infrastructure/wrikeapi"
	"time-track-api/repository"
	"time-track-api/value"
)

var (
	httpClient ch.HttpClient
	timeLogAPI repository.TimeLogRepository
)

func init() {
	httpClient = ch.NewHttpCustomClient(&http.Client{})
	timeLogAPI = wrikeapi.NewTimeLogAPI(httpClient)
}

func main() {
	today := time.Now()
	timeLogs, err := timeLogAPI.FetchByDate(value.NewDate(&today))
	if err != nil {
		log.Fatalln(err)
	}
	for _, timeLog := range timeLogs {
		log.Printf("%+v\n", timeLog)
	}
}
