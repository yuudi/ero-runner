package runner

import (
	"context"
	"fmt"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/filters"
)

func getUserContainer(ctx context.Context, userID string) (string, error) {
	containers, err := docker.ContainerList(ctx, container.ListOptions{
		Filters: filters.NewArgs(filters.Arg("label", "ero-runner.runner.owner="+userID)),
	})
	if err != nil {
		return "", err
	}
	if len(containers) == 0 {
		return "", fmt.Errorf("no container found for user %s", userID)
	}
	return containers[0].ID, nil
}

func createUserContainer(ctx context.Context, userID string) (string, error) {
	resp, err := docker.ContainerCreate(ctx, &container.Config{
		Image: runnerImageID,
		Labels: map[string]string{
			"ero-runner.runner.owner": userID,
		},
	}, nil, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("failed to create container for user %s: %w", userID, err)
	}
	return resp.ID, nil
}
