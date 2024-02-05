package mcndockerclient

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// DockerClient creates a docker client for a given host.
func DockerClient(dockerHost DockerHost) (*client.Client, error) {
	url, err := dockerHost.URL()
	if err != nil {
		return nil, err
	}

	//tlsConfig, err := cert.ReadTLSConfig(url, dockerHost.AuthOptions())
	//if err != nil {
	//	return nil, fmt.Errorf("Unable to read TLS config: %s", err)
	//}

	return client.NewClientWithOpts(client.WithHost(url),
		client.WithTLSClientConfig(dockerHost.AuthOptions().CaCertPath,
			dockerHost.AuthOptions().ClientCertPath,
			dockerHost.AuthOptions().ClientKeyPath))
}

// CreateContainer creates a docker container.
func CreateContainer(dockerHost DockerHost, config *container.Config, containerHostConfig *container.HostConfig, name string) error {
	dockerCli, err := DockerClient(dockerHost)
	if err != nil {
		return err
	}

	if _, err = dockerCli.ImagePull(context.TODO(), config.Image, types.ImagePullOptions{All: true}); err != nil {
		return fmt.Errorf("Unable to pull image: %s", err)
	}

	//var containerConfig container.Config
	//var containerHostConfig container.HostConfig
	var containerNetworkingConfig network.NetworkingConfig
	var platform v1.Platform
	containerResponse, err := dockerCli.ContainerCreate(context.TODO(), config, containerHostConfig, &containerNetworkingConfig, &platform, name)
	if err != nil {
		return fmt.Errorf("Error while creating container: %s", err)
	}

	if err = dockerCli.ContainerStart(context.TODO(), containerResponse.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("Error while starting container: %s", err)
	}

	return nil
}
