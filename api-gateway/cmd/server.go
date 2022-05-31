package main

import "go-micro.dev/v4"

func main() {
	service := micro.NewService(
		micro.Name("hello world"),
	)

	service.Init()

	service.Server()

	service.Run()
}
