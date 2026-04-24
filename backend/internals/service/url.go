package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
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
	customerID, ok := ctx.Value("customer_id").(int)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}

	url := &domain.Url{
		CustomerID:  customerID,
		OriginalUrl: req.OriginalUrl,
	}

	const maxRetries = 5
	for i := 0; i < maxRetries; i++ {
		shortCode, err := shortcode.Generate(url.OriginalUrl)
		if err != nil {
			return nil, err
		}

		url.ShortCode = shortCode

		_, err = s.Repo.InsertUrlKey(ctx, url)
		if err == nil {
			return &dto.UrlShortenResponse{ID: url.ID, ShortCode: shortCode}, nil
		}

		if !isDuplicateError(err) {
			return nil, err
		}
	}

	return nil, fmt.Errorf("failed to generate unique shortcode")
}

func isDuplicateError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}

	return strings.Contains(strings.ToLower(err.Error()), "duplicate key")
}






func (s *UrlService) GetUrl(ctx context.Context, req dto.GetUrlRequest) (*dto.GetUrlResponse, error) {

	bgCtx := context.WithoutCancel(ctx)

	ref := req.Referer
	if ref == "" {
		ref = "direct"
	} else {
		u, err := url.Parse(ref)
		if err != nil || u.Host == "" {
			ref = "unknown"
		} else {
			host := strings.ToLower(u.Host)
			host = strings.TrimPrefix(host, "www.")
			ref = host
		}
	}

	click := domain.ClickEvent{
		UrlID:     0, 
		Referrer:  ref,
		UserAgent: req.UserAgent,
		IP:        req.IP,
	}

	

	url, err := s.Repo.GetUrlFromCache(ctx, req.ShortCode)
	if err == nil {
		click.UrlID = url.ID
		go func() {
			if err := s.Repo.InsertClickEvent(bgCtx, &click); err != nil {
				fmt.Printf("failed to generate click event: %v\n", err)
			}
		}()
		return &dto.GetUrlResponse{
			OriginalUrl: url.OriginalUrl,
		}, nil
	}


	url, err = s.Repo.GetUrlByKey(ctx, req.ShortCode)
	if err != nil {
		return nil, err
	}

	click.UrlID = url.ID
	
	if err := s.Repo.SetUrlInCache(ctx, req.ShortCode, url); err != nil {
		fmt.Printf("failed to set cache: %v\n", err)
	}
	
	go func() {
		if err := s.Repo.InsertClickEvent(bgCtx, &click); err != nil {
			fmt.Printf("failed to generate click event: %v\n", err)
		}
	}()
	

	


	return &dto.GetUrlResponse{
		OriginalUrl: url.OriginalUrl,
	}, nil
}

func (s *UrlService) ListUserURLs(ctx context.Context) ([]dto.ListUrlResponse, error) {
	customerID, ok := ctx.Value("customer_id").(int)
	if !ok {
		return nil, errors.New("unauthorized")
	}

	var urls []dto.ListUrlResponse
	dbUrls, err := s.Repo.ListUserURLs(ctx, customerID)
	if err != nil {
		return nil, err
	}

	for _, dbUrl := range dbUrls {
		urls = append(urls, dto.ListUrlResponse{
			ID:          dbUrl.ID,
			ShortCode:   dbUrl.ShortCode,
			OriginalUrl: dbUrl.OriginalUrl,
		})
	}
	return urls, nil
}
