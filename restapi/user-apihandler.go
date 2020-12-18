package restapi

import (
	"fmt"
	"net/http"
	"strings"
	"zapood/entities"

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
	if operation != "add" && operation != "edit" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	switch strings.ToLower(operation) {
	case "add":
		result := userModel.DB.Create(&user)
		if result.Error != nil {
			fmt.Println("Insert Error :", result.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	case "edit":
		result := userModel.DB.Save(&user)
		if result.Error != nil {
			fmt.Println("Insert Error :", result.Error)
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
		fmt.Println("Delete Error :", result.Error)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	return
}
