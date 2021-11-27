package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/williamwq/go/src/api/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	InitRouter1(router)
	router.GET("/user/users", Users)
	router.GET("/test/tests", Tests)

	router.POST("/user", Store)
	router.GET("/users/:id", Getuser)
	router.PUT("/user/:id", Update)
	router.DELETE("/user/:id", Destroy)
	//用户注册
	router.POST("/user/register", UserRegister)
	//router.POST("/redis", InsertRedis)
	router.POST("/login", Login)

	return router
}
