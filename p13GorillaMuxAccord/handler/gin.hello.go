package handler
import (
	"net/http"

	"github.com/gin-gonic/gin"
)




func (H *DatabaseCollections)Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,"Hello from GIN !!! ")
	}}