package main

import (
	"assets-management/app/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	handler.SetupRoutes(r)

	err := r.Run()
	if err != nil {
		return
	}
}
