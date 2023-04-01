package main

import "github.com/gin-gonic/gin"

func strapiUploads(c *gin.Context) {
	c.Redirect(301, "http://localhost:8082/uploads/"+c.Params[0].Value)
}
