package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harish1907/zgo/controllers"
	"github.com/harish1907/zgo/initializers"
)

func init() {
	initializers.ConnectionDatabase()
	// initializers.MigrationTable()
}

func main() {
	r := gin.Default()
	r.GET("get/all/states", controllers.GetAllStateAPI)
	r.POST("all/village", controllers.GetAllVillageByStateAPI)
	r.POST("add/new/user", controllers.AddNewUserAPI)
	r.Run()
}
