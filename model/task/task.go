package task

import "time"

type Task struct {
	Id           int       `json:"idTask"`
	Title        string    `json:"title"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	Description  string    `json:"description"`
	Rank         int       `json:"rank"`
	PreviousTask int       `json:"prevTask"`
	IdGroup      int       `json:"idGroup"`
}
