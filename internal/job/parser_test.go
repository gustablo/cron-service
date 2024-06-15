package job

import (
	"testing"
	"time"
)

func TestNextExecution(t *testing.T) {
	currentTime := time.Now()

	modifiedTime := time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		0,
		0,
		currentTime.Location(),
	)
	var tests = []struct {
		name  string
		input string
		want  time.Time
	}{
		{"should return a time.Time", "* * * * *", modifiedTime.Add(1 * time.Minute)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NextExecution(tt.input)
			if r != tt.want {
				t.Errorf("got %s, want %s", r, tt.want)
			}
		})
	}
}
