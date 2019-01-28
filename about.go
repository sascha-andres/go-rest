package go_rest

import "github.com/gin-gonic/gin"

func (s *Server) handleAbout() gin.HandlerFunc {
	// you could do a preparation step here which is called only once
	return func(ctx *gin.Context) {
		if _, err := ctx.Writer.WriteString("About"); err != nil {
			_ = ctx.AbortWithError(500, err)
		}
		ctx.Done()
	}
}
