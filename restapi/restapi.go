package restapi

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunApi(DB *gorm.DB) *gin.Engine {
	r := gin.Default()
	RunApiOnRouter(r, DB)
	return r
}

func RunApiOnRouter(r *gin.Engine, DB *gorm.DB) {

	ManageAuthRoutes(r, DB)

	ManageUserRoutes(r, DB)
}

func ManageAuthRoutes(r *gin.Engine, DB *gorm.DB) {
	Handler := NewAuthModel(DB)
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("login", Handler.Login)
	}
}

func ManageUserRoutes(r *gin.Engine, DB *gorm.DB) {
	Handler := NewUserModel(DB)
	//	UserRoutes.Use(Authentication) // middleware fuction runs before api actions
	userGroup := r.Group("/api/User").Use(Authentication())
	{
		userGroup.GET("GetAllUsers", Handler.GetAllUsers)
		userGroup.POST(":operation", Handler.Operation)
		userGroup.DELETE("delete/:id", Handler.Delete)
	}
}
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if info, logedin := IsLogedin(c); logedin {
			// We found the token in our map
			log.Printf("Authenticated user %v\n", info)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}

}
