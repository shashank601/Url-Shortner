package domain

import "time"

type Customer struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
