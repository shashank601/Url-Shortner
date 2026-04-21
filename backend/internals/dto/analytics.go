package dto

type AnalyticsRequest struct {
	CustomerID int    `json:"customer_id"`
	ShortCode  string `json:"short_code"`
}


type AnalyticsResponse struct {
    TotalClicks  int `json:"total_clicks"`
    UniqueClicks int `json:"unique_clicks"`

	// in future, i may add a way to get these analytics
    Browsers  map[string]int `json:"browsers"`
    OS        map[string]int `json:"os"`
    Referrers map[string]int `json:"referrers"`
}
