package main

import (
	"github.com/api-server/controller"
	_ "github.com/api-server/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/auth")
	{
		v1.POST("/loginWithCoulacare", c.ShowResult)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}
