package webapi

import (
	"github.com/gin-gonic/gin"
)

func Boot() {
	gin.SetMode("release")
	router := gin.Default()
	router.Use(corsHandler())

	getAllRouters(router)

	router.RunTLS(":6456", "./resource/ssl/fullchain.pem", "./resource/ssl/privkey.pem")
}
