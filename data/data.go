package data

import "time"

type Task struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Task) FormattedTime() string {
	return t.CreatedAt.Format("2006-01-02 15:04")
}
