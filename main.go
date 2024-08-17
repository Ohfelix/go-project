package main

import (
	"example/hello/config"
	"example/hello/router"

	"fmt"
)

func main() {

	// Initialize config
	err := config.Init()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	router.Initialize()
}
