package apiserver

import "github.com/gin-gonic/gin"

func StringCont(c *gin.Context) {
	firstString := c.Query("firstString")
	secondString := c.Query("secondString")

	c.JSON(200, gin.H{
		"firstString":  firstString,
		"secondString": secondString,
		"result":       firstString + secondString,
	})
}
