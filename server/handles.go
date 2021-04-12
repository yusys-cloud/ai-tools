// Author: yangzq80@gmail.com
// Date: 2021-03-25
//
package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (s *Server) ConfigHandles(r *gin.Engine) {
	rg := r.Group("/api/kv")
	rg.POST("/:b/:k", s.create)
	rg.GET("/:b/:k", s.readAll)
	rg.GET("/:b/:k/:kid", s.read)
	rg.PUT("/:b/:k/:kid", s.update)
	rg.DELETE("/:b/:k/:kid", s.delete)
	rg.DELETE("/:b/:k", s.deleteAll)
	//http
	hg := r.Group("/api/http")
	hg.POST("/do", s.doHttp)
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

func (s *Server) read(c *gin.Context) {

	kv := s.db.Read(c.Param("b"), c.Param("kid"))

	c.JSON(http.StatusOK, kv)
}

func (s *Server) delete(c *gin.Context) {

	s.db.Delete(c.Param("b"), c.Param("kid"))

	c.JSON(http.StatusOK, "ok")
}
func (s *Server) deleteAll(c *gin.Context) {

	i := s.db.DeleteAll(c.Param("b"), c.Param("k"))

	c.JSON(http.StatusOK, "Delete nums:"+strconv.Itoa(i))
}
