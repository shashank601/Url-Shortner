package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shashank601/url-shortner/backend/internals/domain"
)

type CustomerRepo struct {
	DB *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) *CustomerRepo {
	return &CustomerRepo{DB: db}
}

func (r *CustomerRepo) CreateCustomer(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	query := `
	INSERT INTO customers (name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id, name, email, created_at
	`
	
	var createdCustomer domain.Customer
	err := r.DB.QueryRow(ctx, query, customer.Name, customer.Email, customer.Password).Scan(
		&createdCustomer.ID,
		&createdCustomer.Name,
		&createdCustomer.Email,
		&createdCustomer.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdCustomer, nil
}



func (r *CustomerRepo) GetCustomerByEmail(ctx context.Context, email string) (*domain.Customer, error) {
	query := `
	SELECT id, name, email, password, created_at FROM customers WHERE email = $1
	`
	
	var customer domain.Customer
	err := r.DB.QueryRow(ctx, query, email).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Password,
		&customer.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}
