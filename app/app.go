package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApp() {
	router.GET("/", "Hello Alarm")

}
