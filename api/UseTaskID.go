package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UseTaskID(c *gin.Context) {
	data := gin.H{"code": 0, "data": gin.H{"loginResult": true, "taskResult": true}, "message": "success", "traceId": "3737337-hk4e-gens-hin-41df41df41df"}
	c.JSON(http.StatusOK, data)
}
