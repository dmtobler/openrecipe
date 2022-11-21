package api

import (
	"github.com/dmtobler/openrecipe/pkg/openrecipe"
	"github.com/gin-gonic/gin"
	"net/http"
)

// recipe slice to seed recipe data.
var recipes = []openrecipe.Recipe{
	{ID: "1", Title: "Chili", Summary: "Stay warm with chili", Servings: 2},
	{ID: "2", Title: "Chicken Soup", Summary: "Soothe your soul", Servings: 5},
	{ID: "3", Title: "BBQ", Summary: "Texas BBQ", Servings: 10},
}

func Run() {
	router := gin.Default()
	router.GET("/recipes", getRecipes)
	router.GET("/recipes/:id", getRecipeByID)
	router.POST("/recipes", postRecipes)

	router.Run("localhost:8080")
}

// getRecipes responds with the list of all recipes as JSON.
func getRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, recipes)
}

// postRecipe adds a recipe from JSON received in the request body.
func postRecipes(c *gin.Context) {
	var newRecipe openrecipe.Recipe

	// Call BindJSON to bind the received JSON to newRecipe.
	if err := c.BindJSON(&newRecipe); err != nil {
		return
	}

	// Add the new recipe to the slice.
	recipes = append(recipes, newRecipe)
	c.IndentedJSON(http.StatusCreated, newRecipe)
}

// getRecipeByID locates the recipe whose ID value matches the id
func getRecipeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of recipes, looking for a recipe whose ID value matches the parameter.
	for _, r := range recipes {
		if r.ID == id {
			c.IndentedJSON(http.StatusOK, r)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found"})
	return
}
