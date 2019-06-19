package cal

import "github.com/gin-gonic/gin"
func Hello()  {

}

func Router()  {
	routers := gin.Default()
	router := routers.Group("/api/v1")
	router.GET("/hello")
}

