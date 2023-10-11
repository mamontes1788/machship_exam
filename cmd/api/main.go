package main

import "machship/internal/core/dependencyinjection"

func main() {
	server := dependencyinjection.InitializeAPIs()
	server.Start()
}
