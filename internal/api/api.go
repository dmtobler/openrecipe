package api

import (
	"github.com/dmtobler/openrecipe/internal/api/recipe"
	"github.com/dmtobler/openrecipe/internal/api/user"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	// recipes
	router.GET("/recipes", recipe.GetRecipes)
	router.GET("/recipes/:id", recipe.GetRecipeByID)
	router.POST("/recipes", recipe.PostRecipes)

	// users
	router.GET("/users", user.GetUsers)
	router.GET("/users/:id", user.GetUserByID)
	router.POST("/users", user.PostUsers)

	router.Run(":8080")
}
