package main

import "main/package/types"

func main() {
	srv := new(types.Server)
	srv.Serve("8080", nil) // the handler is nil just for a while
}
