package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/williamwq/go/src/api/apis"
)

func InitRouter1(engine *gin.Engine) {
	//router := gin.Default()

	/*router.GET("/users", Users)

	router.POST("/user", Store)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)*/
	engine.POST("/redis", InsertRedis)

	//return router
}
