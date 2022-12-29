package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harish1907/zgo/initializers"
	"github.com/harish1907/zgo/models"
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
	var village []models.Village
	initializers.DB.Find(&village, "state_id=?", body.ID)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"message":     "Data getting successfully.",
		"result":      village,
	})
}

func AddNewUserAPI(c *gin.Context){
	var user models.MyUser
	c.ShouldBind(&user)
	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": 404,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"message": "Create user successfully.",
		"user": user,
	})
}
