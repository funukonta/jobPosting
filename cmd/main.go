package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"redikru/internal/routers"
	"redikru/pkg"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := pkg.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	routers.Routes(mux, db)

	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	fmt.Println("App running, Port", port)
	http.ListenAndServe(port, mux)
}

// docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
