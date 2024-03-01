package main

import (
	"CoffeMapper/app/internal/app"
	"context"
	"fmt"
	"log"
	"time"
)

//Создать CRUD для таблицы с пользователей
//GET /api/v1/users - получает всех пользователей
//GET /api/v1/users/:id - получает пользователя по id
//POST /api/v1/users - создает пользователя
//PUT /api/v1/users/:id - редактирует данные пользователя по id
//DELETE /api/v1/users/:id - удаляет пользователя по id

// GET /api/v1/coffee - получает всех кофейнь
// GET /api/v1/coffee/:id - получает кофейню по id
// POST /api/v1/coffee - создает кофейню
// PUT /api/v1/coffee/:id - редактирует данные кофейни по id
// DELETE /api/v1/coffee/:id - удаляет кофейню по id
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//инициализация БД
	//db, err := postgres.NewPostgresDB()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//инициализация приложения
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal("NewApp error")
	}

	fmt.Println("Run server...")
	err = a.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
