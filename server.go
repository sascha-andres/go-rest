package go_rest

import (
	"fmt"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

// Server struct contains long lived instances
type Server struct {
	router *gin.Engine // http router
	port   int         // port to bind to
	listen string      // device to bind to
}

// NewServer creates ann instance of rest api server
func NewServer(listen string, port int) (*Server, error) {

	r := gin.Default()
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	r.Use(p.Instrument())

	s := &Server{
		router: r,
		port:   port,
		listen: listen,
	}
	s.routes()
	return s, nil
}

// ServeHTTP creates a HTTP listener
func (s *Server) ServeHTTP() error {
	return s.router.Run(fmt.Sprintf("%s:%d", s.listen, s.port))
}

// ServeHTTPS creates a HTTPs listener for a domain with Let's Encrypt
func (s *Server) ServeHTTPS(domain string) error {
	return autotls.Run(s.router, domain)
}

// ServeHTTPS creates a HTTPs listener for a domain with Let's Encrypt
func (s *Server) ServeHTTPSWithCert(certFile, certKey string) error {
	return s.router.RunTLS(fmt.Sprintf("%s:%d", s.listen, s.port), certFile, certKey)
}
