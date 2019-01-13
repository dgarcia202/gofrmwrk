package framework

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/hoisie/web"
)

// requestHandler method that handles a specific type of request
type routeHandler struct {
	pattern string
	method  string
	handler func(*http.Request) http.Response
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
func Route(pattern string, method string, handler func(*http.Request) http.Response) {
	routes = append(routes, routeHandler{pattern, method, handler})
}

func router(ctx *web.Context, path string) {
	log(fmt.Sprintf("received request in path %s", ctx.Request.URL.Path))
	log(fmt.Sprintf("received request in path %s", ctx.Request.RequestURI))

	for _, r := range routes {
		match, err := regexp.MatchString(r.pattern, ctx.Request.URL.Path)
		if err != nil {
			logerr(err)
		} else if match {
			_ = r.handler(ctx.Request)
			ctx.WriteString("something!")
		}
	}

}

func faviconHandler() string {
	return ""
}

func bootstrapApp() {
	web.Get("/favicon.ico", faviconHandler)
	web.Get("/(.*)", router)
	web.Post("/(.*)", router)
}
