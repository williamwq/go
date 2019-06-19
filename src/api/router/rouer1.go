package router

import (
	. "api/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter1(engine *gin.Engine){
	//router := gin.Default()

	/*router.GET("/users", Users)

	router.POST("/user", Store)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)*/
	engine.POST("/redis", InsertRedis)

	//return router
}