package main

import "github.com/gin-gonic/gin"

func main() {
	// Limpar o set mode
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Run(":8080")
}
