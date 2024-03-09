package routes

import (
	"CoffeMapper/app/api/controller"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func InitRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()
	ctx := context.Background()

	users := router.Group("/users")
	{
		users.GET("/", func(c *gin.Context) {
			users, err := controller.GetAllUsers(ctx, db)
			if err != nil {
				return
			}

			c.JSON(200, users)
		})
		users.GET("/:id", func(c *gin.Context) {
			userID := c.Param("id")
			user, err := controller.GetUserById(ctx, db, userID)
			if err != nil {
				c.JSON(404, gin.H{"user not found": "nope"})
				return
			}
			c.JSON(200, user)
		})
		users.POST("/", func(c *gin.Context) {
			var u controller.User
			err := c.BindJSON(&u)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
				return
			}

			err = controller.CreateUser(db, &u)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"id": u.UserId})
		})
		users.PUT("/:id", func(c *gin.Context) {
			userID := c.Param("id")
			var user *controller.User
			err := c.BindJSON(&user)
			fmt.Println(user)
			if err != nil {
				return
			}

			//user не нужен
			newUser, err := controller.UpdateUser(db, user, userID)
			c.JSON(200, gin.H{"newUser": newUser})
		})
		users.DELETE("/:id", func(c *gin.Context) {
			userID := c.Param("id")
			err := controller.DeleteUser(db, userID)
			if err != nil {
				return
			}
			c.Status(http.StatusNoContent)
		})
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
