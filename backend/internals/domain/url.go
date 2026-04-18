package domain

import "time"

type Url struct {
	ID          int
	CustomerID  int
	ShortCode   string
	OriginalUrl string
	CreatedAt   time.Time
	IsActive    bool
}

