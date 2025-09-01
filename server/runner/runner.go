package runner

import (
	"context"
	"log"
	"sync"

	"github.com/moby/moby/api/types/filters"
	"github.com/moby/moby/api/types/image"
	"github.com/moby/moby/client"
)

var docker *client.Client
var runnerImageID string

func initDocker() {
	var err error
	docker, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal("failed to initialize Docker client:", err)
	}
	// check runner image existing
	dockerImages, err := docker.ImageList(context.Background(), image.ListOptions{
		Filters: filters.NewArgs(filters.Arg("label", "ero-runner.runner.base")),
	})
	if err != nil {
		log.Fatal("failed to list Docker images:", err)
	}
	if len(dockerImages) == 0 {
		// todo: build image
		log.Fatal("No runner image found and image build is not implemented yet")
	}
	runnerImageID = dockerImages[0].ID
}

var once sync.Once

func GetDockerClient() *client.Client {
	once.Do(initDocker)
	return docker
}
