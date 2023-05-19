package pleaco

type Container struct {
	Image   string `json:"image"`
	Tag     string `json:"tag"`
	Status  string `json:"status"`
	HasNode bool   `json:"hasNode"`
}

var Containers []Container
