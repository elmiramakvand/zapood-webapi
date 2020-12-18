package restapi

import (
	"fmt"
	"net/http"
	"strings"
	"zapood/entities"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		DB: db,
	}
}

func (userModel UserModel) GetAllUsers(c *gin.Context) {
	var User []entities.User
	rows := userModel.DB.Find(&User)
	sqlRows, _ := rows.Rows()
	var users []entities.User
	for sqlRows.Next() {
		u := entities.User{}
		err2 := sqlRows.Scan(&u.ID, &u.Name, &u.Family, &u.UserName, &u.Password)
		if err2 != nil {
			fmt.Println(err2)
			return
		} else {
			users = append(users, u)
		}
	}
	c.JSON(http.StatusOK, users)
	return
}

func (userModel UserModel) Operation(c *gin.Context) {
	operation := c.Params.ByName("operation")
	if operation == nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Fprintln(w, "operation not found!")
		return
	}
	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Fprintln(w, "could not decode request body by error : %v", err)
		return
	}

	switch strings.ToLower(operation) {
	case "add":
		result := userModel.DB.Create(&user)
		if result.Error != nil {
			fmt.Fprintln(w, "Insert Error : %v", result.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	case "edit":
		result := userModel.DB.Save(&user)
		if result.Error != nil {
			fmt.Fprintln(w, "Insert Error : %v", result.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	return

}

func (userModel UserModel) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var User []entities.User
	result := userModel.DB.Delete(&User, id)
	if result.Error != nil {
		fmt.Fprintln(w, "Delete Error : %v", result.Error)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	return
}

func CheckJWTToken(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("JWTToken")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(err)
		return false
	}
	tokenString := cookie.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(toke *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("token is invalid")
		return false
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("err: %v", err)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("err: %v", err)
		return false
	}
	return true
}
