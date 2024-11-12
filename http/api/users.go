package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roonglit/gin_new/http/views"
)

func (s *Server) GreetUser(c *gin.Context) {
	var req views.GreetUserRequest

	// Bind the request body to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	message := "Hello " + req.Name + "! You are " + strconv.Itoa(req.Age) + " years old."

	res := views.GreetUserResponse{
		Message: message,
	}

	// Return the response
	c.JSON(http.StatusOK, res)
}
