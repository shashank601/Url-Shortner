package service

import (
	"context"
	

	"github.com/shashank601/url-shortner/backend/internals/dto"
	"github.com/shashank601/url-shortner/backend/internals/domain"
	"github.com/shashank601/url-shortner/backend/internals/repo"
	"github.com/shashank601/url-shortner/backend/internals/shortcode"
)

type UrlService struct {
	Repo *repo.UrlRepo
}

func NewUrlService(r *repo.UrlRepo) *UrlService {
	return &UrlService{Repo: r}
}

func (s *UrlService) ShortenUrl(ctx context.Context, req dto.UrlShortenRequest) (*dto.UrlShortenResponse, error) {
	url := &domain.Url{
		CustomerID:  req.CustomerID,
		OriginalUrl: req.OriginalUrl,
	}

	shortCode, err := shortcode.Generate(url.OriginalUrl)
	if err != nil {
		return nil, err
	}

	url.ShortCode = shortCode

	_, err = s.Repo.InsertUrlKey(ctx, url)
	if err != nil {
		return nil, err
	}

	return &dto.UrlShortenResponse{
		ShortCode: shortCode,
	}, nil
}
