package domain

type Analytics struct {
	TotalClicks  int            `json:"total_clicks"`
	UniqueClicks int            `json:"unique_clicks"`
	Browsers     map[string]int `json:"browsers"`
	OS           map[string]int `json:"os"`
	Referrers    map[string]int `json:"referrers"`
}
