# Build stage
FROM golang:1.21-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./
COPY server/ ./server/
COPY todo/ ./todo/

RUN go mod download
RUN go build -o ./bin/server htmx-go-todo/server

# Final stage
FROM scratch

COPY --from=build /build/bin/server /server
COPY templates/ ./templates/
COPY static/ ./static/

EXPOSE 4000

CMD ["/server"]
