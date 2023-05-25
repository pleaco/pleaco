package pleaco

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"pleaco/types"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/containers", getRunningContainersAPI)
	router.POST("/run", runContainerAPI)
	return router
}

func getRunningContainersAPI(c *gin.Context) {
	var runningContainers []pleaco.Container

	for _, runcontainer := range runningContainers {
		if runcontainer.Status == "running" {
			runningContainers = append(runningContainers, runcontainer)
		}
	}

	c.IndentedJSON(http.StatusOK, runningContainers)
}

func runContainerAPI(c *gin.Context) {

	var newContainer pleaco.Container

	err := c.BindJSON(&newContainer)
	if err != nil {
		log.Error(err)
		//c.JSON(http.StatusInternalServerError, "")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Anything except "running" is unexpected here, and we return a 405
	if newContainer.Status != "running" {
		log.Debug("Received unexpected status")
		c.JSON(http.StatusMethodNotAllowed, "Only 'running' allowed as status")
		return
	}

	// Return 201 and the container that was created
	pleaco.Containers = append(pleaco.Containers)
	c.IndentedJSON(http.StatusCreated, newContainer)

}
