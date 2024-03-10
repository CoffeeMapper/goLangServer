package app

import (
	"CoffeMapper/app/api/routes"
	"CoffeMapper/app/pkg/postgres"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

type App struct {
	router     *gin.Engine
	httpServer *http.Server
	pgClient   *sql.DB
}

func NewApp(ctx context.Context) (*App, error) {
	//Раскомментирую позже, нужный функционал для загрузки переменных окружения.
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatal("NewPostgresDB")
	}

	routers := routes.InitRoutes(db)

	return &App{
		pgClient: db,
		router:   routers,
	}, nil
}

// Run method
func (a *App) Run(ctx context.Context) error {
	// /api/v1/
	listener, err := net.Listen("tcp", ":4040")
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
