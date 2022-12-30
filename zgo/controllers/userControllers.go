package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harish1907/zgo/initializers"
	"github.com/harish1907/zgo/models"
	"golang.org/x/crypto/bcrypt"
)

func GetAllStateAPI(c *gin.Context) {
	var states []models.State
	initializers.DB.Select("id", "name").Find(&states)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"message":     "Data getting successfully.",
		"result":      states,
	})
}

func GetAllVillageByStateAPI(c *gin.Context) {
	var body struct {
		ID uint `json:"id"`
	}
	c.Bind(&body)
	var state models.State
	if err := initializers.DB.Preload("Village").Find(&state, body.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": 404,
			"message":     err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"message":     "Data getting successfully.",
		"result":      state,
	})
}

func AddNewUserAPI(c *gin.Context) {
	var user models.MyUser
	c.ShouldBind(&user)
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)
	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": 404,
			"message":     "This user already exist.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"message":     "Create user successfully.",
	})
}

func LoginUserAPI(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	c.BindJSON(&body)
	var findUser models.MyUser
	initializers.DB.First(&findUser, "email=?", body.Email)
	if findUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": 404,
			"message":     "User doesn't exist.",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": 404,
			"message":     "Incorrect password.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"message" : "Login successfully.",
	})
}

func AuthorAwardAPI(c *gin.Context){
	var result models.Award
	initializers.DB.Preload("Author").First(&result, 1)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"result": result,
	})
}
