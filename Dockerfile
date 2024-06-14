# Stage 1: Build the Go binary
FROM golang:1.22.4-alpine

WORKDIR /redikru

COPY . .

# Download dependencies
RUN go mod download

#build
RUN go build -v -o /redikru/web-app ./cmd/main.go

#run app
ENTRYPOINT [ "/redikru/web-app" ]