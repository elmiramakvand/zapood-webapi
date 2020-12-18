package restapi

import (
	"fmt"
	"net/http"
	"time"
	"zapood/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthModel struct {
	DB *gorm.DB
}

var jwtKey = []byte("Zapood_WebApi_2020")

func NewAuthModel(db *gorm.DB) *AuthModel {
	return &AuthModel{
		DB: db,
	}
}

type LoginInfo struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
}

type Claims struct {
	UserName string
	jwt.StandardClaims
}

func (authModel AuthModel) Login(c *gin.Context) {
	var loginInfo LoginInfo
	err := c.BindJSON(&loginInfo)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var User []entities.User
	result := authModel.DB.Where("UserName = ? AND Password >= ?", loginInfo.UserName, loginInfo.Password).Find(&User)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	sqlRow := result.Row()
	u := entities.User{}
	err2 := sqlRow.Scan(&u.ID, &u.Name, &u.Family, &u.UserName, &u.Password)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	expireTime := time.Now().Add(8 * time.Hour)
	claims := &Claims{
		UserName: u.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(jwtKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, stringToken)
	return
}

func IsLogedin(c *gin.Context) (LoginInfo, bool) {
	tokenString := c.GetHeader("Authorization")
	if len(tokenString) == 0 {
		return LoginInfo{}, false
	}
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		return LoginInfo{}, false
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return LoginInfo{}, false
		}
		return LoginInfo{}, false
	}

	return LoginInfo{
		UserName: claims.UserName,
	}, true
}
