package utils

import "github.com/gin-gonic/gin"

func ServerErrorResponse(Error error) gin.H {
	if gin.Mode() == "debug" {
		return gin.H{"status": "server faile", "Error": Error}
	} else {
		return gin.H{"status": "server faile"}

	}
}
