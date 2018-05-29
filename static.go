package go_rest

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleStatic(path, prefix string, indexes bool) gin.HandlerFunc {
	// you could do a preparation step here which is called only once
	return static.Serve(prefix, static.LocalFile(path, indexes))
}
