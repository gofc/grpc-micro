package registry

//Service config info
type Service struct {
	Name    string  `json:"name"`
	Version string  `json:"version"`
	Nodes   []*Node `json:"nodes"`
}

//Node config info
type Node struct {
	Id      string `json:"id"`
	Address string `json:"address"`
}
