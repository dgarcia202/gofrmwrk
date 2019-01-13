package main

import (
	"net/http"

	app "github.com/dgarcia202/ancestry/framework"
)

func main() {
	app.Route("/test", "GET", func(*http.Request) http.Response {
		return http.Response{StatusCode: 201}
	})
	app.Serve()
}
