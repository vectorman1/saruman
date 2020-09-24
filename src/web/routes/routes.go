package routes

import (
	"github.com/vectorman1/saruman/src/consts"
	"github.com/vectorman1/saruman/src/web/handlers"
	"net/http"
)

var GETRoutesMap = make(map[string]func(http.ResponseWriter, *http.Request))
var POSTRoutesMap = make(map[string]func(w http.ResponseWriter, r *http.Request))

var GetRouters = [...]string {
	consts.IndexUrl,
	consts.ConfigAppUrl,
}

var PostRouters = [...]string {
	consts.ConfigBaseUrl,
}

var GetHandlers = [...]func(http.ResponseWriter, *http.Request) {
	handlers.IndexHandler,
	handlers.ConfigAppGetHandler,
}

var PostHandlers = [...]func(w http.ResponseWriter, r *http.Request) {
	handlers.ConfigPostHandler,
}

func InitializeMap() {
	for i, route := range GetRouters {
		GETRoutesMap[route] = GetHandlers[i]
	}
	for i, route := range PostRouters {
		POSTRoutesMap[route] = PostHandlers[i]
	}
}