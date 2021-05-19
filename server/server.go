// Author: yangzq80@gmail.com
// Date: 2021-03-16
//
package server

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/yusys-cloud/ai-tools/conf"
	"github.com/yusys-cloud/ai-tools/server/web"
	"github.com/yusys-cloud/go-jsonstore-rest/rest"
)

type Server struct {
	cf *conf.Conf
}

func NewServer(cf *conf.Conf) *Server {

	return &Server{cf}
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
		//config := cors.DefaultConfig()
		//config.AllowAllOrigins = true
		//engine.Use(cors.New(config))
	}

	engine.Use(static.Serve("/", static.LocalFile("./ui", false)))

	s.ConfigHandles(engine)

	rest.NewJsonStoreRest(s.cf.Path, engine)

	ConfigUserHandles(engine)

	engine.Run(":" + s.cf.Port)
}
