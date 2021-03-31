// Author: yangzq80@gmail.com
// Date: 2021-03-16
//
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yusys-cloud/ai-tools/conf"
	"github.com/yusys-cloud/ai-tools/server/db"
	"github.com/yusys-cloud/ai-tools/server/web"
)

type Server struct {
	db *db.Storage
	cf *conf.Conf
}

func NewServer(cf *conf.Conf) *Server {

	return &Server{db.NewStorage(cf.Path), cf}
}

func (s *Server) Start() {
	s.startApiServer()
}

func (s *Server) startApiServer() {
	if s.cf.Mode == "pro" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()

	//Needed in order to disable CORS for local development
	if s.cf.Mode == "dev" {
		engine.Use(web.DisableCors())
	}

	s.ConfigHandles(engine)

	engine.Run(":" + s.cf.Port)
}
