package app

import (
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/api/helper"
	"log"
	"net/http"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "Hospital API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var input helper.LoginInput

		if err := c.ShouldBindJSON(&input); err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse(nil))
			return
		}

		token, err := s.authService.Login(input)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse(token))
	}
}
