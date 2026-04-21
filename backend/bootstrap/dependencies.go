package bootstrap

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/shashank601/url-shortner/backend/internals/handler"
	"github.com/shashank601/url-shortner/backend/internals/repo"
	"github.com/shashank601/url-shortner/backend/internals/service"
)

type Dependencies struct {
	UrlHandler       *handler.UrlHandler
	AuthHandler      *handler.AuthHandler
	AnalyticsHandler *handler.AnalyticsHandler
}

func InitDependencies(db *pgxpool.Pool, rdb *redis.Client) *Dependencies {

	urlRepo := repo.NewUrlRepo(db, rdb)
	customerRepo := repo.NewCustomerRepo(db)
	analyticsRepo := repo.NewAnalyticsRepo(db)

	urlService := service.NewUrlService(urlRepo)
	authService := service.NewAuthService(customerRepo)
	analyticsService := service.NewAnalyticsService(analyticsRepo, urlRepo)

	urlHandler := handler.NewUrlHandler(urlService)
	authHandler := handler.NewAuthHandler(authService)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsService)

	return &Dependencies{
		UrlHandler:       urlHandler,
		AuthHandler:      authHandler,
		AnalyticsHandler: analyticsHandler,
	}
}
