package http

import (
	"net/http"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) createUser(c *gin.Context) {
	data := models.User{}
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.userService.Create(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}
