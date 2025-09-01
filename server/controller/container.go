package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuudi/ero-runner/server/model"
	"github.com/yuudi/ero-runner/server/runner"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func RunCommand(c *gin.Context) {
	userValue, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, ok := userValue.(model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
		return
	}

	command := c.PostForm("command")
	if command == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Command is required"})
		return
	}

	err := runner.RunCommand(c, user, command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute command"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Command executed", "user": user})
}
