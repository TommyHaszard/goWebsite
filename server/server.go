package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
	"time"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
	"github.com/patrickmn/go-cache"
)

type ErrorHandlerFunc func(err error, c *Context)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Ctx      context.Context
	Cache    *cache.Cache
}

func (c *Context) Render(component templ.Component) error {
	return component.Render(c.Ctx, c.Response)
}

func (c *Context) FormValue(name string) string {
	return c.Request.FormValue(name)
}

type MiddlewareFunc func(HandlerFunc) HandlerFunc

type HandlerFunc func(c *Context) error

type Server struct {
	router       *httprouter.Router
	ErrorHandler ErrorHandlerFunc
	middlewares  []MiddlewareFunc
	cache        *cache.Cache
}

func New() *Server {
	return &Server{
		router:       httprouter.New(),
		ErrorHandler: defaultErrorHandler,
		cache:        cache.New(5*time.Minute, 10*time.Minute),
	}
}
func (s *Server) Start(port string) error {
	return http.ListenAndServe(port, s.router)
}

func (s *Server) Get(path string, h HandlerFunc, middleware ...MiddlewareFunc) {
	s.router.GET(path, s.makeHTTPRouterHandle(h, middleware...))
}

func (s *Server) Post(path string, h HandlerFunc, middleware ...MiddlewareFunc) {
	s.router.POST(path, s.makeHTTPRouterHandle(h, middleware...))
}

func (s *Server) makeHTTPRouterHandle(h HandlerFunc, mws ...MiddlewareFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ctx := &Context{
			Response: w,
			Request:  r,
			Ctx:      context.Background(),
			Cache:    s.cache,
		}
		handlerWithMiddleware := applyMiddleware(h, mws...)
		if err := handlerWithMiddleware(ctx); err != nil {
			s.ErrorHandler(err, ctx)
		}
	}
}

func (s *Server) InsertMiddleware(middlewares ...MiddlewareFunc) {
	s.middlewares = append(s.middlewares, middlewares...)
}

func defaultErrorHandler(err error, c *Context) {
	pc, _, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	formatted := fmt.Sprintf("%s: %d |", funcName, line)
	slog.Error(formatted, "ERROR", err)
}

// func (c *Context) GetResponse() http.ResponseWriter {
// 	return c.response
// }

// func (c *Context) GetRequest() *http.Request {
// 	return c.request
// }

func (c *Context) GetContext(key string) any {
	return c.Ctx.Value(key)
}

func (c *Context) SetContext(key string, value any) {
	c.Ctx = context.WithValue(c.Ctx, key, value)
}

// func (c *Context) GetCache() *cache.Cache {
// 	return c.cache
// }

func applyMiddleware(h HandlerFunc, mw ...MiddlewareFunc) HandlerFunc {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}
	return h
}
