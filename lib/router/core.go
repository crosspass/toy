package router

import (
	"net/http"
)

type Handler interface {
	Show(*http.Request) ([]byte, error)
	Create(*http.Request) ([]byte, error)
	Update(*http.Request) ([]byte, error)
	Edit(*http.Request) ([]byte, error)
	Destroy(*http.Request) ([]byte, error)
}

type HandlerFunc func(*http.Request) ([]byte, error)

type router struct {
	Path            string
	Handler         Handler
	RootHanlderFunc HandlerFunc
}

var Routers = make(map[string]router)

func Resources(path string, handler Handler) {
	Routers[path] = router{Path: path, Handler: handler}
}

func Root(handlerFunc HandlerFunc) {
	Routers["/"] = router{Path: "/", RootHanlderFunc: handlerFunc}
}
