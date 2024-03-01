package app

import (
	"CoffeMapper/app/api/routes"
	"CoffeMapper/app/pkg/postgres"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
)

type App struct {
	router     *gin.Engine
	httpServer *http.Server
	pgClient   *postgres.PostgresDB
}

func NewApp(ctx context.Context) (*App, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPostgresDB()

	routers := routes.InitRoutes(db)

	return &App{
		pgClient: db,
		router:   routers,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", ":4040/api/v1/")
	if err != nil {
		log.Fatalf("Run error: %v", err)
	}

	handler := a.router

	a.httpServer = &http.Server{
		Handler: handler,
	}

	err = a.httpServer.Serve(listener)
	if err != nil {
		log.Fatalf("Servak ypal: %v", err)
	}

	return err
}
