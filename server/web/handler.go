// Author: yangzq80@gmail.com
// Date: 2021-03-31
//
package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Needed in order to disable CORS for local development
func DisableCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
