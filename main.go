package main

import (
	"main/package/handler"
	"main/package/types"
)

func main() {
	srv := new(types.Server)
	handler := new(handler.Handler)
	srv.Serve("8080", handler.ServeRoutes())
}
