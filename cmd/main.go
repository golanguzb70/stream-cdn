package main

import (
	"fmt"

	"github.com/golanguzb70/stream-cdn/internal/proxy"
)

func main() {
	server := proxy.NewProxy()

	fmt.Printf("Starting server on port %d\n", server.Config.Port)
	err := server.Start()
	if err != nil {
		fmt.Println(err)
	}
}
