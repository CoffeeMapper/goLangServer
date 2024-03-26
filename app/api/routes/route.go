package routes

import (
	"CoffeMapper/app/api/controller"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Создать CRUD для таблицы с брэндами
//
//GET /api/v1/brand- получает все брэнды
//GET /api/v1/brand/:id - получает брэнды по id
//POST /api/v1/brand- создает брэнды
//PUT /api/v1/brand/:id - редактирует данные брэнда по id
//DELETE /api/v1/brand/:id - удаляет брэнды по id

//Создать CRUD для таблицы с пользователей
//GET /api/v1/users - получает всех пользователей
//GET /api/v1/users/:id - получает пользователя по id
//POST /api/v1/users - создает пользователя
//PUT /api/v1/users/:id - редактирует данные пользователя по id
//DELETE /api/v1/users/:id - удаляет пользователя по id

func InitRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()
	ctx := context.Background()

	//CRUD -
	//CREATE - POST
	//READ - GET
	//UPDATE - PUT
	//DELETE - DELETE
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
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

		coffee := api.Group("/coffee")
		{
			coffee.GET("/")
			coffee.GET("/:id")
			coffee.POST("/")
			coffee.PUT("/:id")
			coffee.DELETE("/:id")
		}

		brand := api.Group("/brand")
		{
			brand.GET("/", func(c *gin.Context) {
				brands, err := controller.GetAllBrands(ctx, db)
				if err != nil {
					return
				}

				c.JSON(200, brands)
			})
			brand.GET("/:id", func(c *gin.Context) {
				brandID := c.Param("id")
				brand, err := controller.GetBrandById(ctx, db, brandID)
				if err != nil {
					c.JSON(404, gin.H{"user not found": "nope"})
					return
				}
				c.JSON(200, brand)
			})
			brand.POST("/", func(c *gin.Context) {
				var b controller.Brand
				err := c.BindJSON(&b)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid brand data"})
					return
				}

				err = controller.CreateBrand(db, &b)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusCreated, gin.H{"id": b.BrandId})
			})
			brand.PUT("/:id", func(c *gin.Context) {
				brandID := c.Param("id")
				var brand *controller.Brand
				err := c.BindJSON(&brand)
				fmt.Println(brand)
				if err != nil {
					return
				}

				//brand не нужен
				newBrand, err := controller.UpdateBrand(db, brand, brandID)
				c.JSON(200, gin.H{"newUser": newBrand})
			})
			brand.DELETE("/:id", func(c *gin.Context) {
				brandID := c.Param("id")
				err := controller.DeleteBrand(db, brandID)
				if err != nil {
					return
				}
				c.Status(http.StatusNoContent)
			})
		}
	}

	return router
}
