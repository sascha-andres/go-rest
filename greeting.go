package go_rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleGreeting(name string) gin.HandlerFunc {
	// you could do a preparation step here which is called only once
	return func(ctx *gin.Context) {
		to := ctx.Param("to")
		if to != "" {
			_, _ = ctx.Writer.WriteString(fmt.Sprintf("Hello from %s to %s", name, to))
		} else {
			_, _ = ctx.Writer.WriteString(fmt.Sprintf("Hello from %s", name))
		}
		ctx.Done()
	}
}
