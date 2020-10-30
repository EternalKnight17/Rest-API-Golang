package main

import (
	"fmt"

	"./apiserver"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Rest API Golang")

	r := gin.Default()
	r.GET("/", apiserver.StringCont)
	r.Run("localhost:5000")
}
