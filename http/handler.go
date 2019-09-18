package http

import (
	"log"
	"todo/config"
	"todo/services"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Server creates a server for exposing application over HTTP protocol
type Server struct {
	config         *config.Config
	taskService    *services.TaskService
	userService    *services.UserService
	authMiddleware *jwt.GinJWTMiddleware
}

// Run creates router and run server
func (s *Server) Run() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", s.authMiddleware.LoginHandler)
		v1.POST("/users", s.createUser)

		v1.Use(s.authMiddleware.MiddlewareFunc())
		{

			v1.GET("/tasks", s.allTasks)
			v1.GET("/tasks/:id", s.getTask)
			v1.POST("/tasks", s.createTask)
			v1.PUT("/tasks/:id", s.updateTask)
		}
	}

	router.Run(":" + s.config.Port)
}

// NewServer Initialize server with its dependencies
func NewServer(config *config.Config, taskService *services.TaskService, userService *services.UserService) *Server {
	authMiddleware, err := getAuthMiddleware(userService)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		config:         config,
		taskService:    taskService,
		userService:    userService,
		authMiddleware: authMiddleware,
	}
}
