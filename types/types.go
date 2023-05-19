package pleaco

type Container struct {
	Image   string `json:"image"`
	Tag     string `json:"tag"`
	Status  string `json:"status"`
	HasNode bool   `json:"hasNode"`
	Name    string `json:"name"`
	Id      string `json:"id"`
}

var Containers []Container
