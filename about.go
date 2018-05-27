package go_rest

import "github.com/gin-gonic/gin"

func (s *Server) handleAbout() gin.HandlerFunc {
	// you could do a preparation step here which is called only once
	return func(ctx *gin.Context) {
		ctx.Writer.WriteString("About")
		ctx.Done()
	}
}
