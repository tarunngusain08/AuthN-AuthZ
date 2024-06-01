package models

import "time"

type User struct {
	ID        int        `db:"id"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt *time.Time `db:"created_at"`
}
