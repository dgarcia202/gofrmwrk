package main

import (
	app "github.com/dgarcia202/ancestry/framework"
)

type customer struct {
	name    string
	address string
}

func main() {

	app.Route("^/test$", "GET", func(request *app.ActionRequest) app.ActionResponse {
		return app.ActionResponse{}
	})
	/*
		app.Route("^/customers$", "POST", func(c *customer) app.ActionResponse {
			return http.Response{StatusCode: 201}
		})

		app.Route("^/customers/(.*)$", "GET", func(string) customer {
			return http.Response{StatusCode: 201}
		})
	*/
	app.Serve()
}
