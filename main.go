package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	// "time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	reader, err := cli.ImagePull(
		context.Background(),
		"nginx",
		types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	showContainerAll(cli)

	createContainer(cli)

	containers := getCountainerAll(cli)
	showContainers(containers)

	for _, container := range containers {
		startContainer(cli, container)
	}

	containers = getCountainerAll(cli)
	showContainers(containers)

	for _, container := range containers {
		stopContainer(cli, container)
	}

	containers = getCountainerAll(cli)
	showContainers(containers)

	for _, container := range containers {
		removeContainer(cli, container)
	}

	showContainerAll(cli)
}

func getCountainerAll(cli *client.Client) []types.Container {
	containers, err := cli.ContainerList(
		context.Background(),
		types.ContainerListOptions{All: true},
	)
	if err != nil {
		panic(err)
	}
	return containers
}

func showContainerAll(cli *client.Client) {
	containers := getCountainerAll(cli)
	showContainers(containers)
}

func showContainers(containers []types.Container) {
	if len(containers) == 0 {
		println("not found containers")
		return
	}

	for _, container := range containers {
		var ports []string
		for _, port := range container.Ports {
			ports = append(ports, fmt.Sprintf("%s:%d", port.IP, port.PublicPort))
		}

		fmt.Printf("%s %s %s %s\n", container.ID[:10], container.Image, container.State, ports)
	}
}

func startContainer(cli *client.Client, container types.Container) {
	if container.State != "exited" && container.State != "created" {
		return
	}

	err := cli.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	println(container.ID + " started")
}

func stopContainer(cli *client.Client, container types.Container) {
	if container.State != "running" {
		return
	}

	timeout := time.Duration(500) * time.Millisecond
	err := cli.ContainerStop(
		context.Background(),
		container.ID,
		&timeout,
	)
	if err != nil {
		panic(err)
	}

	println(container.ID + " stopped")
}

func createContainer(cli *client.Client) {
	result, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: "nginx",
		}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}
	println(result.ID + " created")
}

func removeContainer(cli *client.Client, container types.Container) {
	if container.State == "exited" {
		err := cli.ContainerRemove(
			context.Background(),
			container.ID,
			types.ContainerRemoveOptions{},
		)
		if err != nil {
			panic(err)
		}

		println(container.ID + " removed")
	}
}
