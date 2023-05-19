package pleaco

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	pleaco "pleaco/types"

	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunContainers() {
	for {
		for _, schedulecontainer := range pleaco.Containers {
			if schedulecontainer.HasNode == false {

				// Lock the container for other nodes
				schedulecontainer.HasNode = true

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

				resp, err := cli.ContainerCreate(ctx, &container.Config{
					Image: schedulecontainer.Image,
				}, nil, nil, nil, "")
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
	}
}
