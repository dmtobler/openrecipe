package user

import (
	"github.com/dmtobler/openrecipe/pkg/openrecipe"
	"github.com/gin-gonic/gin"
	"net/http"
)

// user slice to seed recipe data.
var users = []openrecipe.User{
	{ID: "1", UserName: "dmtobler", NickName: "Danielle"},
	{ID: "2", UserName: "jltobler", NickName: "Justin"},
}

// GetUsers responds with the list of all users as JSON.
func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// PostUsers adds a user from JSON received in the request body.
func PostUsers(c *gin.Context) {
	var newUser openrecipe.User

	// Call BindJSON to bind the received JSON to newUser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new user to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// GetUserByID locates the user whose ID value matches the id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of users, looking for a user whose ID value matches the parameter.
	// TODO: Refactor for efficiency
	for _, r := range users {
		if r.ID == id {
			c.IndentedJSON(http.StatusOK, r)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	return
}
