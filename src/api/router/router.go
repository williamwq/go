package router

import (
	"github.com/gin-gonic/gin"
	. "api/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	InitRouter1(router)
	router.GET("/users", Users)

	router.POST("/user", Store)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)
	//router.POST("/redis", InsertRedis)

	return router
}