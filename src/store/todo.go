package store

import "time"

type Todo struct {
	Id          string
	Title       string `json:"title"`
	Description string `json:"description"`
	Label       string `json:"label"`
	Deadline    string `json:"deadline"`
	Severity    string `json:"severity"`
	Priority    string `json:"priority"`
	Created_at  time.Time
	Deleted_at  time.Time
}
