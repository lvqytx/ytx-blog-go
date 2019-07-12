package router

import (
	"net/http"
	"ytx-blog-go/server"
)

type Router struct {
}

func (p *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		server.SayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
