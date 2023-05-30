package pleaco

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	pleaco "pleaco/types"
	"time"

	"context"
	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunContainers() {
	for {
		for _, schedulecontainer := range pleaco.Containers {
			if schedulecontainer.HasNode == false {

				// Lock the container for other nodes
				schedulecontainer.HasNode = true

				// Add heartbeat timestamp
				schedulecontainer.Heartbeat = time.Now().Unix()

				// Fail if we can't get our own hostname
				hostname, err := os.Hostname()
				if err != nil {
					log.Fatal(err)
				}

				// Add hostname as `Node`
				schedulecontainer.Node = hostname

				ctx := context.Background()
				cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
				if err != nil {
					log.Error(err)
					schedulecontainer.HasNode = false
					continue
				}
				defer cli.Close()

				out, err := cli.ImagePull(ctx, schedulecontainer.Image, types.ImagePullOptions{})
				if err != nil {
					log.Error(err)
					schedulecontainer.HasNode = false
					continue
				}
				defer out.Close()
				io.Copy(os.Stdout, out)

				resp, err := cli.ContainerCreate(ctx, &containertypes.Config{
					Image: schedulecontainer.Image,
				}, nil, nil, nil, schedulecontainer.Name)
				if err != nil {
					log.Error(err)
					schedulecontainer.HasNode = false
					continue
				}

				if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
					log.Error(err)
					schedulecontainer.HasNode = false
					continue
				}

				log.Debug(resp.ID)

			}

			// Determine age of heartbeat of containers and try to schedule them here if outdated
			heartbeatAge := time.Now().Unix() - schedulecontainer.Heartbeat
			if heartbeatAge >= 120 {

				// Lock the container for other nodes
				schedulecontainer.HasNode = true
				schedulecontainer.Heartbeat = time.Now().Unix()
				hostname, err := os.Hostname()
				if err != nil {
					log.Fatal(err)
				}
				schedulecontainer.Node = hostname

				ctx := context.Background()
				cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
				if err != nil {
					log.Error(err)
					schedulecontainer.HasNode = false
					continue
				}
				defer cli.Close()

				out, err := cli.ImagePull(ctx, schedulecontainer.Image, types.ImagePullOptions{})
				if err != nil {
					log.Error(err)
					schedulecontainer.HasNode = false
					continue
				}
				defer out.Close()
				io.Copy(os.Stdout, out)

				resp, err := cli.ContainerCreate(ctx, &containertypes.Config{
					Image: schedulecontainer.Image,
				}, nil, nil, nil, schedulecontainer.Name)
				if err != nil {
					log.Error(err)
					schedulecontainer.HasNode = false
					continue
				}

				if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
					log.Error(err)
					schedulecontainer.HasNode = false
					continue
				}

				log.Debug(resp.ID)

			}
		}
		// Sleep before next iteration
		time.Sleep(10 * time.Second)
	}
}

func DeleteContainers() {
	for {
		hostname, err := os.Hostname()
		if err != nil {
			log.Fatal(err)
			continue
		}

		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Error(err)
		}
		defer cli.Close()

		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			log.Error(err)
		}

		// Check all running containers against the list of containers we should be running
		for _, container := range containers {
			log.Debug(container.ID)
			shouldRun := false
			for _, toBeRunning := range pleaco.Containers {
				if toBeRunning.Node == hostname {
					shouldRun = true
				}
			}
			noWaitTimeout := 0
			if shouldRun == false {
				err := cli.ContainerStop(ctx, container.ID, containertypes.StopOptions{Timeout: &noWaitTimeout})
				if err != nil {
					log.Error(err)
				}
			}

		}
		// Sleep before next iteration
		time.Sleep(10 * time.Second)
	}
}
