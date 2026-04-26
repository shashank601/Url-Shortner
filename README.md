
## URL Shortner
 
### Prerequisites
 
- Go (matching `backend/go.mod`)
- Node.js + npm (for the `client/`)
- Postgres
- Redis
 
### Environment variables
 
Backend reads env vars from `.env` (loaded at startup).
 
- `PORT` (optional, default `8080`)
- `DATABASE_URL` (required)
- `REDIS_URL` (required)
- `JWT_SECRET` (required)
 
Example:
 
```env
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/url_shortner?sslmode=disable
REDIS_URL=redis://localhost:6379/0
JWT_SECRET=change_me
```
 
### Run backend
 
From `backend/`:
 
- Run migrations (use whatever tool you already use for `backend/migrations/`)
- Start server
 
The API will run on `http://localhost:8080`.
 
### Run client
 
From `client/`:
 
- `npm install`
- `npm run dev`
 
Client defaults to calling `http://localhost:8080`.
 
To override:
 
- Set `VITE_API_URL` in `client/.env`.
 
Example:
 
```env
VITE_API_URL=http://localhost:8080
```
 
### API routes (backend)
 
- `POST /signup`
- `POST /login`
- `GET /verify` (requires `Authorization: Bearer <token>`)
- `POST /shorten` (requires auth)
- `GET /urls` (requires auth)
- `GET /analytics/{code}` (requires auth)
- `GET /{code}` (public redirect)
 

 