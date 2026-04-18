package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	
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

	createdCustomer, err := s.Repo.CreateCustomer(ctx, customer)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		ID:    createdCustomer.ID,
		Name:  createdCustomer.Name,
		Email: createdCustomer.Email,
	}, nil
}




func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	customer, err := s.Repo.GetCustomerByEmail(ctx, req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password))
	if err != nil {
		return dto.LoginResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    customer.ID,
		"name":  customer.Name,
		"email": customer.Email,
	})

	tokenString, err := token.SignedString([]byte("JWT_SECRET_KEY"))
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
		Token: tokenString,
	}, nil
}

