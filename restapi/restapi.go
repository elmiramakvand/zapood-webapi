package restapi

import (
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

// func ManageAuthRoutes(r *gin.Engine, DB *gorm.DB) {
// 	Handler := NewAuthModel(DB)
// 	AuthRoute := r.PathPrefix("/api/auth").Subrouter()
// 	AuthRoute.Methods("POST").Path("/login").HandlerFunc(Handler.Login)
// }

func ManageUserRoutes(r *gin.Engine, DB *gorm.DB) {
	Handler := NewUserModel(DB)
	//	UserRoutes.Use(Authentication) // middleware fuction runs before api actions
	userGroup := r.Group("/api/User")
	{
		userGroup.GET("GetAllUsers", Handler.GetAllUsers)
		userGroup.POST("{operation:(?:add|edit)}", Handler.Operation)
		userGroup.DELETE("delete/:id", Handler.Delete)
	}
}

// func Authentication(next http.Handler) http.Handler {

// 	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if info, logedin := IsLogedin(r); logedin {
// 			// We found the token in our map
// 			log.Printf("Authenticated user %v\n", info)
// 			next.ServeHTTP(w, r)
// 		} else {
// 			http.Error(w, "شما به این صفحه دسترسی ندارید", http.StatusForbidden)
// 		}
// 	})

// 	return h
// }
