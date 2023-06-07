package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/api"
	"log"
	"os"
)

type Server struct {
	router               *gin.Engine
	authService          api.AuthService
	roleService          api.RoleService
	userService          api.UserService
	permissionService    api.PermissionService
	hspDisciplineService api.HspDisciplineService
	hspUnitService       api.HspUnitService
	hspDoctorService     api.HspDoctorService
}

func NewServer(
	router *gin.Engine,
	authService api.AuthService,
	roleService api.RoleService,
	userService api.UserService,
	permissionService api.PermissionService,
	hspDisciplineService api.HspDisciplineService,
	hspUnitService api.HspUnitService,
	hspDoctorService api.HspDoctorService,
) *Server {
	return &Server{
		router:               router,
		authService:          authService,
		roleService:          roleService,
		userService:          userService,
		permissionService:    permissionService,
		hspDisciplineService: hspDisciplineService,
		hspUnitService:       hspUnitService,
		hspDoctorService:     hspDoctorService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	server := fmt.Sprintf(":%s", port)
	err := r.Run(server)

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
