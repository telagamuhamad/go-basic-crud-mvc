package controllers

import (
	"fmt"
	"net/http"
	"testing-app/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	user := models.User{Name: name, Email: email}
	result := models.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := models.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	fmt.Println(users)

	c.HTML(http.StatusOK, "user_list.gohtml", gin.H{"users": users})
}

func ShowUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "user_form.gohtml", nil)
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("ID Received:", id)

	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	c.HTML(http.StatusOK, "user_detail.gohtml", gin.H{"user": user})
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	c.HTML(http.StatusOK, "user_edit_form.gohtml", gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	name := c.PostForm("name")
	email := c.PostForm("email")

	user.Name = name
	user.Email = email

	result := models.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	result := models.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}
