package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"github.com/shashank601/url-shortner/backend/internals/domain"
	"github.com/shashank601/url-shortner/backend/internals/dto"
	"github.com/shashank601/url-shortner/backend/internals/repo"
)


type AuthService struct {
	Repo *repo.CustomerRepo
}


func NewAuthService(r *repo.CustomerRepo) *AuthService {
	return &AuthService{Repo: r}
}


func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {


	customer := &domain.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)


	if err != nil {
		return dto.RegisterResponse{}, err
	}

	customer.Password = string(hashed)

	id, name, email, err := s.Repo.CreateCustomer(ctx, customer)

	response := dto.RegisterResponse{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return response, err
}