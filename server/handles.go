// Author: yangzq80@gmail.com
// Date: 2021-03-25
//
package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) ConfigHandles(r *gin.Engine) {
	rg := r.Group("/api")
	rg.GET("/kv/:b", s.getKeys)
	rg.GET("/kv/:b/:k", s.getOne)
	rg.POST("/kv/:b/:k", s.save)
}

func (s *Server) save(c *gin.Context) {

	v := c.DefaultPostForm("v", "{}")

	s.db.Save(c.Param("b"), c.Param("k"), v)

}

func (s *Server) getKeys(c *gin.Context) {
	b := s.db.GetAll(c.Param("b"))
	c.JSON(http.StatusOK, b)
}

func (s *Server) getOne(c *gin.Context) {

	b := s.db.GetOne(c.Param("b"), c.Param("k"))

	c.JSON(http.StatusOK, b)
}
