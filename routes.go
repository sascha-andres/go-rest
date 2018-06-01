package go_rest

import "os"

func (s *Server) routes() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	if stat, err := os.Stat("templates"); err == nil && stat.IsDir() {
		s.router.LoadHTMLGlob("templates/*")
	}
	s.router.Handle("GET", "/about/", s.handleAbout())
	s.router.Handle("GET", "/greet/*to", s.handleGreeting(name))
	s.router.Handle("GET", "/template/", s.handleTemplate())
	s.router.Use(s.handleStatic("./static", "/", true))
}

/*

Middleware attaching:

s.router.Handle("GET", "/admin", s.adminOnly(s.handleAdminIndex))

func (s *Server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if !currentUser(r).IsAdmin {
            http.NotFound(w, r)
            return
        }
        h(w, r)
    }
}

*/
