// Author: yangzq80@gmail.com
// Date: 2021-03-25
//
package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) ConfigHandles(r *gin.Engine) {
	rg := r.Group("/api/kv/:b/:k")
	rg.POST("/", s.create)
	rg.GET("/", s.readAll)
	rg.GET("/:kid", s.readOne)
	rg.PUT("/:kid", s.update)
	rg.DELETE("/:kid", s.delete)
	rg.DELETE("/", s.deleteAll)
}

func (s *Server) create(c *gin.Context) {

	var data interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := s.db.Create(c.Param("b"), c.Param("k"), data)

	c.JSON(http.StatusOK, id)
}

func (s *Server) update(c *gin.Context) {
	var data interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.Update(c.Param("b"), c.Param("kid"), data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (s *Server) readAll(c *gin.Context) {
	b := s.db.ReadAll(c.Param("b"), c.Param("k"))
	c.JSON(http.StatusOK, b)
}

func (s *Server) readOne(c *gin.Context) {

	kv := s.db.ReadOne(c.Param("b"), c.Param("kid"))

	c.JSON(http.StatusOK, kv)
}

func (s *Server) delete(c *gin.Context) {

	s.db.Delete(c.Param("b"), c.Param("kid"))

	c.JSON(http.StatusOK, "ok")
}
func (s *Server) deleteAll(c *gin.Context) {
	b := c.Param("b")
	rs := s.db.ReadAll(b, c.Param("k"))
	for _, value := range rs {
		s.db.Delete(b, value.K)
	}
	c.JSON(http.StatusOK, "ok")
}
