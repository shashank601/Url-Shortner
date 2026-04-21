package dto

type UrlShortenRequest struct {
	OriginalUrl string `json:"original_url"`
}

type UrlShortenResponse struct {
	ID        int    `json:"url_id"`
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
