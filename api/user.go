package api

import (
	"example/web3-profile-api/db"
	"example/web3-profile-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := userByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
func userByID(id string) (*model.User, error) {
	for i, u := range db.UsersTable {
		if u.ID == id {
			return &db.UsersTable[i], nil
		}
	}
	return nil, errors.New("user not found")
}
