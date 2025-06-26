package task

import "time"

type TaskFilter struct {
	Status        string // "new", "in_progress", "done"
	AuthorID      int   
	CreatedAfter  time.Time
	CreatedBefore time.Time
	UpdatedAfter  time.Time
	UpdatedBefore time.Time
}
