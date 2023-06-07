package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"hospital-api/pkg/api/helper"
	"hospital-api/pkg/repository/model"
	"net/http"
)

func GatePermission(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var permission model.CorePermission
		var role model.CoreRole
		var permissions []model.CorePermission
		jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
				Message: err.Error(),
			})
			return
		}

		token, err := parseToken(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
				Message: "bad jwt token",
			})
			return
		}

		claims, OK := token.Claims.(jwt.MapClaims)
		if !OK {
			c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
				Message: "unable to parse claims",
			})
			return
		}

		for _, data := range claims["permission"].([]interface{}) {
			marshal, _ := json.Marshal(data)
			_ = json.Unmarshal(marshal, &permission)
			permissions = append(permissions, permission)
		}

		for _, data := range claims["role"].([]interface{}) {
			marshal, _ := json.Marshal(data)
			_ = json.Unmarshal(marshal, &role)
			for _, item := range role.Permission {
				marshalItem, _ := json.Marshal(item)
				_ = json.Unmarshal(marshalItem, &permission)
				permissions = append(permissions, permission)
			}
		}

		for _, data := range helper.Unique(permissions) {
			if data.Name == scope {
				fmt.Println("Check Roles passed")
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "no access provided",
		})
		return
	}
}
