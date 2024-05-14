package util

import "time"

type Time interface {
	Now() time.Time
}

type timer struct{}

func NewTimer() Time {
	return timer{}
}

func (timer) Now() time.Time {
	return time.Now()
}
