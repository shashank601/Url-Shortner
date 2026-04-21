package service

import (
	"context"
	"fmt"

	"github.com/shashank601/url-shortner/backend/internals/dto"
	"github.com/shashank601/url-shortner/backend/internals/repo"
)

type AnalyticsService struct {
	AnalyticsRepo *repo.AnalyticsRepo
	UrlRepo     *repo.UrlRepo
}

func NewAnalyticsService(analyticsRepo *repo.AnalyticsRepo, urlRepo *repo.UrlRepo) *AnalyticsService {
	return &AnalyticsService{
		AnalyticsRepo: analyticsRepo,
		UrlRepo:     urlRepo,
	}
}


func (s *AnalyticsService) GetAnalytics(ctx context.Context, req dto.AnalyticsRequest) (dto.AnalyticsResponse, error) {
	
	url, err := s.UrlRepo.GetUrlByKey(ctx, req.ShortCode)
	if err != nil {
		return dto.AnalyticsResponse{}, err
	}

	
	if url.CustomerID != req.CustomerID {
		return dto.AnalyticsResponse{}, fmt.Errorf("unauthorized: URL does not belong to customer")
	}

	analytics, err := s.AnalyticsRepo.GetAnalytics(ctx, req.CustomerID, url.ID)
	if err != nil {
		return dto.AnalyticsResponse{}, err
	}

	return dto.AnalyticsResponse{
		TotalClicks:  analytics.TotalClicks,
		UniqueClicks: analytics.UniqueClicks,
		Browsers:     analytics.Browsers,
		OS:           analytics.OS,
		Referrers:    analytics.Referrers,
	}, nil
}
