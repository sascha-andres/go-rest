package go_rest

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func (s *Server) handleGreeting(name string) gin.HandlerFunc {
	// you could do a preparation step here which is called only once
	return func(ctx *gin.Context) {
		ctx.Writer.WriteString(fmt.Sprintf("Hello from %s", name))
		ctx.Done()
	}
}