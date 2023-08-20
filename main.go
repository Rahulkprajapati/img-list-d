package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	// Create a new Docker client
	// The client provides a Go API for interacting with Docker
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// List all images
    // The ImageList method returns a list of images on the Docker host
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	// Print image details
	// The image details are stored in the Image struct
	fmt.Println("Docker images on local system:")
	for _, image := range images {
		fmt.Printf("ID: %s, RepoTags: %v\n", image.ID, image.RepoTags)
	}
}