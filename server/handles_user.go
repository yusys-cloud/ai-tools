// Author: yangzq80@gmail.com
// Date: 2021-04-06
//
package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigUserHandles(r *gin.Engine) {
	rg := r.Group("/vue-element-admin")
	rg.POST("/user/login", login)
	rg.GET("/user/info", info)
}

func login(c *gin.Context) {

	jsonData := []byte(`{"code":20000,"data":{"token":"admin-token"}}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)

	c.JSON(http.StatusOK, v)
}

func info(c *gin.Context) {

	jsonData := []byte(`{"code":20000,"data":{"roles":["admin"],"introduction":"I am a super administrator","avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif","name":"Super Admin"}}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)

	c.JSON(http.StatusOK, v)
}
