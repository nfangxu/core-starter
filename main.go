package main

import (
	"github.com/DoNewsCode/core-starter/internal/bootstrap"
)

func main() {
	// Core Bootstrap
	root, shutdown := bootstrap.Bootstrap()

	// Shutdown
	defer shutdown()

	// Run
	_ = root.Execute()
}
