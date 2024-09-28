package model

import "time"

type Secret struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	Type      string    `db:"type"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
