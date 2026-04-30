FROM node:20-alpine AS frontend-builder
WORKDIR /app/client
COPY client/package*.json ./
RUN npm ci
COPY client/ ./
RUN npm run build



FROM golang:1.26-alpine AS backend-builder
WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .


COPY --from=frontend-builder /app/client/dist ./static 
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /url-shortener ./cmd/app/main.go


FROM alpine:3.20
RUN apk --no-cache add ca-certificates

RUN addgroup -S app && adduser -S app -G app
WORKDIR /app

COPY --from=backend-builder /url-shortener ./url-shortener
COPY --from=backend-builder /app/static ./static


USER app

EXPOSE 8080
CMD ["./url-shortener"]