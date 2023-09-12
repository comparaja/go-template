package domain

import "time"

type Todo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Todo) ChangeTitle(title string) error {
	t.Title = title

	return nil
}
