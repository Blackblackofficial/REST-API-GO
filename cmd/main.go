package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/pgxpool"
	"github.com/joho/godotenv"
	"REST-API/internal/persons/delivery"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	connection, state := os.LookupEnv("DATABASE_URL")
	if !state {
		log.Print("connection string was not found")
		fmt.Print()
	}

	conn, err := pgxpool.Connect(context.Background(), connection)
	if err != nil {
		log.Fatal("database connection not established")
	}

	port, ok := os.LookupEnv("PORT")

	if ok == false {
		port = "5000"
	}

	pr := repo.NewPRepository(*conn)
	pu := usecase.NewPUsecase(pr)
	//pd := delivery.NewPHandler(pu)

	r := mux.NewRouter()
	r.Use(middleware.InternalServerError)

}