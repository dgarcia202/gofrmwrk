package framework

import (
	"github.com/hoisie/web"
)

// requestHandler method that handles a specific type of request
type routeHandler struct {
	pattern string
	method  string
	handler interface{}
}

// ActionRequest contains relevant information about the incoming request
type ActionRequest struct {
	params map[string]string
	body   interface{}
}

// ActionResponse allows request handlers to customize the generate http response
type ActionResponse struct {
	status int
	body   interface{}
}

var routes = make([]routeHandler, 0)

// Serve starts the request listener
func Serve() {
	ServeOn("0.0.0.0:9000")
}

// ServeOn starts the application listening to a specific net interface and port
func ServeOn(netInterface string) {
	bootstrapApp()
	web.Run(netInterface)
}

// Route adds a new unbinded route handler
func Route(pattern string, method string, handler interface{}) {
	routes = append(routes, routeHandler{pattern, method, handler})
}

func faviconHandler() string {
	return ""
}

func bootstrapApp() {
	web.Get("/favicon.ico", faviconHandler)
	web.Get("/(.*)", router)
	web.Post("/(.*)", router)
}
