package main

import (
	route "api-go/router"
	"fmt"
)

func main() {
	fmt.Print("Server Loaded")
	route.StartServer().Run(":8086")
}
