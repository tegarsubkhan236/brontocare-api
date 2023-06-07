package app

import (
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/api/helper"
	"hospital-api/pkg/repository/model"
	"log"
	"strconv"
)

func (s *Server) permissionDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse("ID not found"))
			return
		}

		data, err := s.permissionService.Detail(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse(data))
	}
}

func (s *Server) ListPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.AbortWithStatusJSON(helper.BadResponse(err))
			return
		}
		perPageStr := c.DefaultQuery("per_page", "10")
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.AbortWithStatusJSON(helper.BadResponse(err))
			return
		}

		roles, err := s.permissionService.List(page, perPage)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse(roles))
	}
}

func (s *Server) CreatePermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newData model.NewCorePermission

		err := c.ShouldBindJSON(&newData)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse(err))
			return
		}

		err = s.permissionService.New(newData)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse("New Role has been created"))
	}
}

func (s *Server) DeletePermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse("ID not found"))
			return
		}

		err = s.permissionService.Delete(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse("failed to update"))
			return
		}

		c.JSON(helper.SuccessResponse(nil))
	}
}
