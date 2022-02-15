package main

import (
	"fmt"
	_ "huangcoolyserver/docs"
	"huangcoolyserver/src/controller"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Demo
// @version 1.0
// @description Swagger API.
// @host localhost:8888
func main() {

	router := gin.Default()
	router.Use(CORSMiddleware())
	c := controller.NewController()

	employeeController := router.Group("/huangcoolyserver/empolyee")
	{
		employeeController.GET("getEmployee", c.GetEmployee)
		employeeController.POST("addEmployee", c.AddEmployee)
		employeeController.PUT("updateEmployee", c.UpdateEmployee)
	}

	userController := router.Group("/huangcoolyserver/user")
	{
		userController.GET("getUser", c.GetUser)
		userController.POST("login", c.Login)
		userController.POST("checkUserLoginStatus", c.CheckUserLoginStatus)
		userController.GET("getRoleAccessSystem/:userId", c.GetRoleAccessSystem)
	}

	pgConnString := fmt.Sprintf("http://localhost:%s/swagger/doc.json",
		os.Getenv("SWAGGERPORT"))
	fmt.Println(pgConnString)
	url := ginSwagger.URL(pgConnString) // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	ginPort := fmt.Sprintf(":%s",
		os.Getenv("GINPORT"))
	fmt.Println(ginPort)
	router.Run(ginPort)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
