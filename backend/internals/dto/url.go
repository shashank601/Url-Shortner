package dto

type UrlShortenRequest struct {
	OriginalUrl string `json:"original_url"`
	CustomerID  int    `json:"customer_id"`
}

type UrlShortenResponse struct {
	ShortCode string `json:"short_code"`
}

type GetUrlRequest struct {
	ShortCode  string `json:"short_code"`
	UserAgent  string `json:"user_agent"`
	Referer    string `json:"referer"`
	IP         string `json:"ip"`
}

type GetUrlResponse struct {
	OriginalUrl string `json:"original_url"`
}
