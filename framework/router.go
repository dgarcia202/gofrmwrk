package framework

import (
	"fmt"
	"regexp"

	"github.com/hoisie/web"
)

func router(ctx *web.Context, path string) {

	log("received request to path %s", ctx.Request.URL.Path)

	routeFound := false
	for _, r := range routes {
		match, err := regexp.MatchString(r.pattern, ctx.Request.URL.Path)
		if err != nil {
			logerr(err)
			break
		} else if match {
			log("matched route with pattern %s", r.pattern)
			executeAction(ctx, &r)
			routeFound = true
			break
		}
	}

	if !routeFound {
		ctx.NotFound(fmt.Sprintf("No defined route matches path '%s'\n", ctx.Request.URL.Path))
	}
}

func executeAction(ctx *web.Context, r *routeHandler) {

	_ = r.handler(ctx.Request)
	ctx.WriteString("something!")
}
