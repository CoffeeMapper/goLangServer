package routes

import (
	"CoffeMapper/app/pkg/postgres"
	"github.com/gin-gonic/gin"
)

//GET /api/v1/coffee - получает всех кофейнь
//GET /api/v1/coffee/:id - получает кофейню по id
//POST /api/v1/coffee - создает кофейню
//PUT /api/v1/coffee/:id - редактирует данные кофейни по id
//DELETE /api/v1/coffee/:id - удаляет кофейню по id

//Создать CRUD для таблицы с пользователей
//GET /api/v1/users - получает всех пользователей
//GET /api/v1/users/:id - получает пользователя по id
//POST /api/v1/users - создает пользователя
//PUT /api/v1/users/:id - редактирует данные пользователя по id
//DELETE /api/v1/users/:id - удаляет пользователя по id

func InitRoutes(db *postgres.PostgresDB) *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	{
		users.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"Jopa": "xer"})
		})
		users.GET("/:id")
		users.POST("/")
		users.PUT("/:id")
		users.DELETE("/:id")
	}

	coffee := router.Group("/coffee")
	{
		coffee.GET("/")
		coffee.GET("/:id")
		coffee.POST("/")
		coffee.PUT("/:id")
		coffee.DELETE("/:id")
	}

	brand := router.Group("/brand")
	{
		brand.GET("/")
		brand.GET("/:id")
		brand.POST("/")
		brand.PUT("/:id")
		brand.DELETE("/:id")
	}

	return router
}
