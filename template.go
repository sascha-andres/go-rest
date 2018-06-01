package go_rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleTemplate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Templated call",
		})
	}
}
