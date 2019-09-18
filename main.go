package main

import (
	"fmt"

	"go.uber.org/dig"

	"todo/config"
	"todo/http"
	"todo/models"
	"todo/services"
	"todo/storage"
)

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.NewDatabase)
	container.Provide(config.NewConfig)
	container.Provide(storage.NewAdapter)
	container.Provide(models.NewTaskRepository)
	container.Provide(models.NewUserRepository)
	container.Provide(services.NewTaskService)
	container.Provide(services.NewUserService)
	container.Provide(http.NewServer)

	return container
}

func main() {
	container := buildContainer()

	err := container.Invoke(func(server *http.Server, config *config.Config) {
		fmt.Println("Listening on http://localhost:" + config.Port)
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
