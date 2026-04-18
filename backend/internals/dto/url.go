package dto

type UrlShortenRequest struct {
	OriginalUrl string `json:"original_url"`
	CustomerID  int    `json:"customer_id"`
}

type UrlShortenResponse struct {
	ShortCode string `json:"short_code"`
}
