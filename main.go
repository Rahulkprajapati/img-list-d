package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// List all images
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	// Print image details
	fmt.Println("Docker images on local system:")
	for _, image := range images {
		fmt.Printf("ID: %s, RepoTags: %v\n", image.ID, image.RepoTags)
	}
}