package main

import (
	"fmt"

	"github.com/go-example/api/gin/examples/context-cancel-request/pkg/http/rest"
)

func main() {
	fmt.Println("Starting context cancel request")

	app := rest.NewApp()

	app.Run(":8000")
}
