package docker

import (
	"context"
	"github.com/bitmyth/accounts/src/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"os"
)

type Client struct {
	ctx         context.Context
	cli         *client.Client
	ContainerID string
}

func New() (*Client, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &Client{ctx, cli, ""}, nil
}

func (c *Client) Images() []types.ImageSummary {
	images, err := c.cli.ImageList(c.ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	return images
}
func (c *Client) IsImagePulled(name string) bool {

	for _, image := range c.Images() {
		for _, tag := range image.RepoTags {
			if tag == name {
				println(name, "exists! skip pulling")
				return true
			}
		}
	}
	return false
}

func (c *Client) IsImageRunning(image string) bool {
	containers, err := c.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		if c.Image == image {
			return true
		}
	}
	return false
}

func (c *Client) RunMySQL() error {
	cli, err := New()
	if err != nil {
		return err
	}
	ctx := cli.ctx

	image := "mysql:5.7"

	if !c.IsImagePulled(image) {
		reader, err := cli.cli.ImagePull(ctx, "docker.io/library/mysql:5.7", types.ImagePullOptions{})
		if err != nil {
			return err
		}
		io.Copy(os.Stdout, reader)
	}

	if c.IsImageRunning(image) {
		println(image, "is running")
		return nil
	}

	hostConfig := &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{
			"3306/tcp": {
				{"localhost", "3306"},
			},
		},
	}

	resp, err := cli.cli.ContainerCreate(ctx, &container.Config{
		Image: image,
		Cmd:   []string{"--character-set-server=utf8mb4", "--default-time-zone=+08:00"},
		Tty:   false,
		Env:   []string{"MYSQL_ROOT_PASSWORD=123", "MYSQL_DATABASE=accounts"},
	}, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	c.ContainerID = resp.ID

	return nil
}

func (c Client) Stop(ID string) error {
	return c.cli.ContainerStop(c.ctx, ID, nil)
}

func (c *Client) RunMigration(dsn string) error {
	cli, err := New()
	if err != nil {
		return err
	}
	ctx := cli.ctx

	image := "bitmyth/goose:latest"

	if !c.IsImagePulled(image) {
		reader, err := cli.cli.ImagePull(ctx, image, types.ImagePullOptions{})
		if err != nil {
			return err
		}
		io.Copy(os.Stdout, reader)
	}

	if c.IsImageRunning(image) {
		println(image, "is running")
	}

	hostConfig := &container.HostConfig{
		Binds: []string{
			config.RootPath + "src/database/migrations:/migrations",
		},
	}
	networkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}

	resp, err := cli.cli.ContainerCreate(ctx, &container.Config{
		Image:      image,
		Cmd:        []string{"goose", "up"},
		Tty:        true,
		Env:        []string{"GOOSE_DRIVER=mysql", "GOOSE_DBSTRING=" + dsn},
		WorkingDir: "/migrations",
	}, hostConfig, networkingConfig, nil, "")

	if err != nil {
		panic(err)
	}

	if err := cli.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	println(resp.ID)

	out, err := c.cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	//go func() {
	//stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	io.Copy(os.Stdout, out)
	//}()

	return nil
}
