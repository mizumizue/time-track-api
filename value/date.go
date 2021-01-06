package value

import "time"

type Date string

func NewDate(date *time.Time) Date {
	return Date(date.Format("2006-01-02"))
}
