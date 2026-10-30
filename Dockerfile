# Frontend build
FROM node:alpine AS app-builder
WORKDIR /app
COPY package.json package-lock.json* ./
RUN npm install
COPY . .
RUN npm run build

# Go server build (pure Go SQLite, no CGO)
FROM golang:alpine AS server-builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go build -o triangle_travel .
RUN go build -o seed ./cmd/seed

# Final image
FROM alpine
WORKDIR /app
COPY --from=app-builder /app/build ./build
COPY --from=server-builder /app/triangle_travel .
COPY --from=server-builder /app/seed .
COPY db/schema.sql db/seed_data.sql ./db/
EXPOSE 8080
RUN ./seed && rm ./seed
ENV ENV=production
# PORT is set by Render; -host 0.0.0.0 for container
CMD ["./triangle_travel", "-host", "0.0.0.0"]
