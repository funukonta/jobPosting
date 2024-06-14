package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"redikru/internal/routers"
	"redikru/pkg"

	_ "redikru/docs"

	_ "github.com/joho/godotenv/autoload"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Swagger API Documentation - Redikru Job Posting
// @description This is a API Documentation for Redikru Job Posting.

// @Tags Jobs
// @Tags Companies

func main() {
	db, err := pkg.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	routers.Routes(mux, db)

	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	fmt.Println("App running, Port", port)
	http.ListenAndServe(port, mux)
}

// docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
