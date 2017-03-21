package servemux

import (
	"net/http"

	http2 "github.com/golang-plus/net/http"

	"github.com/julienschmidt/httprouter"
)

// New returns a new Router.
func New() http2.ServeMux {
	return &ServeMux{
		Router: httprouter.New(),
	}
}

type ServeMux struct {
	*httprouter.Router
}

func (s *ServeMux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.Router.ServeHTTP(rw, req)
}

func (s *ServeMux) wrap(handler http.Handler) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		r2 := http2.NewRequest(r)
		for _, v := range params {
			r2.SetParameter(v.Key, v.Value)
		}
		handler.ServeHTTP(rw, r2.Request)
	}
}

func (s *ServeMux) Handle(method string, path string, handler http.Handler) {
	s.Router.Handle(method, path, s.wrap(handler))
}
