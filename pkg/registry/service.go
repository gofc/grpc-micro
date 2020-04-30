package registry

//Service config info
type Service struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	//Version string `json:"version"`
	Address string `json:"address"`
}
