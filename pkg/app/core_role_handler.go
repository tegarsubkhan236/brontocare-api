package app

import (
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/api/helper"
	"hospital-api/pkg/repository/model"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) CreateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newData model.NewCoreRole

		err := c.ShouldBindJSON(&newData)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse(err))
			return
		}

		err = s.roleService.New(newData)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse("New Role has been created"))
	}
}

func (s *Server) roleDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse("ID not found"))
			return
		}

		data, err := s.roleService.Detail(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse(data))
	}
}

func (s *Server) ListRole() gin.HandlerFunc {
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

		roles, err := s.roleService.List(page, perPage)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(helper.SuccessResponse(roles))
	}
}

func (s *Server) AssignRolePermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var Data model.AssignPermissionToRole

		//id := uuid.Must(uuid.FromString(c.Param("id")))
		id, _ := strconv.Atoi(c.Param("id"))

		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.AbortWithStatusJSON(helper.BadResponse(err))
			return
		}

		err = s.roleService.AssignPermission(id, Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.AbortWithStatusJSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse("Permission has been updated"))
	}
}

func (s *Server) UpdateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var Data model.CoreRole
		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse("can't bind the value"))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse("ID not found"))
			return
		}

		err = s.roleService.Update(id, Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse("failed to update"))
			return
		}

		c.JSON(helper.SuccessResponse(nil))
	}
}

func (s *Server) DeleteRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse("ID not found"))
			return
		}

		err = s.roleService.Delete(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse("failed to update"))
			return
		}

		c.JSON(helper.SuccessResponse(nil))
	}
}

func (s *Server) BatchDeleteRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var Data model.BatchDeleteRole
		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse(err))
			return
		}

		err = s.roleService.BatchDelete(Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse(nil))
	}
}
