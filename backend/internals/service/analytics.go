package service

import (
	"context"

	"github.com/shashank601/url-shortner/backend/internals/dto"
	"github.com/shashank601/url-shortner/backend/internals/repo"
)

type AnalyticsService struct {
	Repo *repo.AnalyticsRepo	
}

func NewAnalyticsService(repo *repo.AnalyticsRepo) *AnalyticsService {
	return &AnalyticsService{
		Repo: repo,
	}
}


func (s *AnalyticsService) GetAnalytics(ctx context.Context, req dto.AnalyticsRequest) (dto.AnalyticsResponse, error) {
	analytics, err := s.Repo.GetAnalytics(ctx, req.CustomerID, req.UrlID)
	if err != nil {
		return dto.AnalyticsResponse{}, err
	}
	
	return dto.AnalyticsResponse{
		TotalClicks: analytics.TotalClicks,
		UniqueClicks: analytics.UniqueClicks,
		Browsers: analytics.Browsers,
		OS: analytics.OS,
		Referrers: analytics.Referrers,
	}, nil
}
