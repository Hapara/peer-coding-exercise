package main

import (
	"fmt"
	"peer-coding/cmd/hapara"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/cache/:id", func(context *gin.Context) {
		id := context.Param("id")

		value := hapara.DoWork(id)

		context.JSON(200, gin.H{
			fmt.Sprintf("%s", id): value,
		})
	})

	router.Run(":8080")
}
