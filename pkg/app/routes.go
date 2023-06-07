package app

import (
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/app/middleware"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router
	router.Use(middleware.CORS())

	v1 := router.Group("/v1/api")
	{
		unProtected := v1.Group("/")
		{
			unProtected.GET("/status", s.ApiStatus())
			unProtected.POST("/login", s.Login())
			unProtected.POST("/register", s.CreateUser())
		}

		protected := v1.Group("/")
		{
			protected.Use(middleware.JwtTokenCheck)
			user := protected.Group("/user").Use(middleware.GatePermission("manage-user"))
			{
				user.POST("/create", s.CreateUser())
				user.GET("/list", s.ListUser())
				user.GET("/detail", s.UserDetail())
				user.PUT("/update/:id", s.UpdateUser())
				user.PUT("/update_password/:id", s.UpdateUserPassword())
				user.PUT("/assign_permission/:id", s.AssignUserPermission())
				user.PUT("/assign_role/:id", s.AssignUserRole())
				user.DELETE("/delete/:id", s.DeleteUser())
			}
			permission := protected.Group("/permission").Use(middleware.GatePermission("manage-permission"))
			{
				permission.GET("/list", s.ListPermission())
				permission.GET("/detail/:id", s.permissionDetail())
				permission.POST("/create", s.CreatePermission())
				permission.DELETE("/delete/:id", s.DeletePermission())
			}
			role := protected.Group("/role").Use(middleware.GatePermission("manage-role"))
			{
				role.POST("/create", s.CreateRole())
				role.GET("/list", s.ListRole())
				role.GET("/detail/:id", s.roleDetail())
				role.PUT("/update/:id", s.UpdateRole())
				role.PUT("/assign_permission/:id", s.AssignRolePermission())
				role.DELETE("/delete/:id", s.DeleteRole())
			}
		}
	}

	return router
}
