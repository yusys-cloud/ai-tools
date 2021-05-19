// Author: yangzq80@gmail.com
// Date: 2021-03-25
//
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yusys-cloud/ai-tools/ssh"
)

func (s *Server) ConfigHandles(r *gin.Engine) {
	//http
	hg := r.Group("/api/http")
	hg.POST("/do", s.doHttp)
	//Websocket
	r.GET("/api/ws", ssh.WsSsh)
}
