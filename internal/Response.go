package internal

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, httpStatus int, response map[string]interface{}) {
	var r gin.H
	r = response
	c.JSON(httpStatus, r)
}
