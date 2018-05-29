package go_rest

import "os"

func (s *Server) routes() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	s.router.Handle("GET", "/about/", s.handleAbout())
	s.router.Handle("GET", "/greet/*to", s.handleGreeting(name))
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
