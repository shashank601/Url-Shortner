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

func (r *CustomerRepo) CreateCustomer(ctx context.Context, customer *domain.Customer) (int, string, string, error) {
	
	query := `
	INSERT INTO customers (name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id, name, email
	`
	
	var id int
	var name string
	var email string
	err := r.DB.QueryRow(ctx, query, customer.Name, customer.Email, customer.Password).Scan(&id, &name, &email)


	if err != nil {
		return 0, "", "", err
	}

	return id, name, email, nil
}


// for setup
// import (
//     "context"
//     "fmt"
//     "os"

//     "github.com/jackc/pgx/v5/pgxpool"
// )

// func main() {
//     ctx := context.Background()
//     connStr := "postgres://username:password@localhost:5432/dbname"

//     // Create the connection pool
//     pool, err := pgxpool.New(ctx, connStr)
//     if err != nil {
//         fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
//         os.Exit(1)
//     }
//     defer pool.Close() // Ensure the pool is closed when the app exits

//     // Verify the connection
//     if err := pool.Ping(ctx); err != nil {
//         fmt.Printf("Ping failed: %v\n", err)
//         return
//     }

//     fmt.Println("Connected to PostgreSQL!")
// }