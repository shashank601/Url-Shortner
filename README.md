# URL Shortener

<div align="center">

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white&labelColor=000)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white&labelColor=000)
![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white&labelColor=000)
![React](https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=white&labelColor=000)
![Vite](https://img.shields.io/badge/Vite-646CFF?style=for-the-badge&logo=vite&logoColor=white&labelColor=000)
![TailwindCSS](https://img.shields.io/badge/TailwindCSS-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white&labelColor=000)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white&labelColor=000)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white&labelColor=000)

</div>

![Screenshot Placeholder](./screenshot-placeholder.png)

A modern URL shortening service with user authentication, analytics, and click tracking. Transform long URLs into short, shareable links while gaining insights into click patterns, user agents, and referrer data.

## Main Feature

**Analytics Dashboard** - Track every click on your shortened URLs with total clicks and unique visitor counts. URL lookups are cached using Redis for high-performance redirects.

## Features

- **User Authentication** - Secure signup and login with JWT-based authentication
- **URL Shortening** - Generate unique short codes for any URL
- **URL Management** - View and manage all your shortened URLs in one place
- **Click Analytics** - Basic analytics including:
  - Total click count
  - Unique click tracking
- **Public Redirect** - Fast redirection from short codes to original URLs
- **Redis Caching** - High-performance caching for URL lookups

## Tech Stack

### Backend
- **Go 1.26.2** - Core backend framework
- **PostgreSQL** - Primary database for users, URLs, and click events
- **Redis** - Caching layer for URL lookups
- **JWT (golang-jwt/jwt)** - Authentication token management
- **pgx/v5** - PostgreSQL driver
- **go-redis/v9** - Redis client

### Frontend
- **React 19.2.5** - UI framework
- **Vite 8.0.10** - Build tool and dev server
- **TailwindCSS 4.2.4** - Styling
- **React Router 7.14.2** - Client-side routing
- **Axios** - HTTP client for API requests

## Prerequisites

- Go 1.26.2 or higher
- Node.js 18+ and npm
- PostgreSQL 12+
- Redis 6+

## Setup Instructions

### 1. Clone the Repository

```bash
git clone <repository-url>
cd Url-Shortner
```

### 2. Backend Setup

Navigate to the backend directory:

```bash
cd backend
```

Create a `.env` file with the following environment variables:

```env
PORT=8080
DATABASE_URL="postgresql://user:password@host:5432/dbname?sslmode=require"
REDIS_URL="rediss://default:token@host:6379"
JWT_SECRET="your_jwt_secret_here"
BASE_URL="http://localhost:8080"
JWT_EXPIRES_IN="24h"
VITE_API_URL="/api"
```

Install Go dependencies:

```bash
go mod download
```

Run database migrations:

```bash
# Use your preferred PostgreSQL migration tool
# Example with migrate:
migrate -path migrations -database "$DATABASE_URL" up
```

Start the backend server:

```bash
go run cmd/app/main.go
```

The API will be available at `http://localhost:8080`

### 3. Frontend Setup

Navigate to the client directory:

```bash
cd client
```

Create a `.env` file with the API URL:

```env
VITE_API_URL=http://localhost:8080
```

Install dependencies:

```bash
npm install
```

Start the development server:

```bash
npm run dev
```

The frontend will be available at `http://localhost:5173`

## API Endpoints

### Authentication

- **POST /signup**
  - Register a new user
  - Body: `{"name": "string", "email": "string", "password": "string"}`

- **POST /login**
  - Login with email and password
  - Body: `{"email": "string", "password": "string"}`
  - Returns JWT token

- **GET /verify**
  - Verify JWT token validity
  - Headers: `Authorization: Bearer <token>`

### URL Management

- **POST /shorten**
  - Create a shortened URL
  - Headers: `Authorization: Bearer <token>`
  - Body: `{"url": "string"}`

- **GET /urls**
  - List all URLs for authenticated user
  - Headers: `Authorization: Bearer <token>`

### Analytics

- **GET /analytics/{code}**
  - Get analytics for a specific short code
  - Headers: `Authorization: Bearer <token>`
  - Returns: `{"total_clicks": int, "unique_clicks": int}`

### Public

- **GET /r/{code}**
  - Redirect to original URL
  - Tracks click event with IP, user agent, and referrer

## Environment Variables

### Backend (.env)

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| PORT | No | 8080 | Server port |
| DATABASE_URL | Yes | - | PostgreSQL connection string |
| REDIS_URL | Yes | - | Redis connection string |
| JWT_SECRET | Yes | - | Secret key for JWT signing |
| BASE_URL | Yes | - | Base URL for shortened links (e.g., http://localhost:8080) |

### Frontend (.env)

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| VITE_API_URL | No | http://localhost:8080 | Backend API URL |

## Future Scope

- [ ] Custom alias support for short codes
- [ ] URL expiration dates
- [ ] QR code generation for shortened URLs
- [ ] Bulk URL shortening
- [ ] Export analytics data (CSV, JSON)
- [ ] API rate limiting
- [ ] Webhook notifications on clicks