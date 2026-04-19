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


type ClickEvent struct {
	ID         int
	UrlID      int
	Referrer   string
	UserAgent  string
	IP         string
	CreatedAt  time.Time
}



