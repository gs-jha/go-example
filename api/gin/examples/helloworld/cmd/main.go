package main

import (
	"fmt"

	"github.com/go-example/api/gin/examples/helloworld/pkg/http/rest"
)

func main() {
	fmt.Println("Starting hello world app")

	// Setup routes
	router := rest.SetRoutes()

	// Run server on port 8000
	router.Run(":8000")

	fmt.Println("Server running on port 8000")
}
