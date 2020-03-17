package registry

type Service struct {
	Name    string  `json:"name"`
	Version string  `json:"version"`
	Nodes   []*Node `json:"nodes"`
}

type Node struct {
	Id      string `json:"id"`
	Address string `json:"address"`
}
