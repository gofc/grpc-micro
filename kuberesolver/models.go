package kuberesolver

//EventType is type of a event
type EventType string

const (
	//Added is a add event
	Added EventType = "ADDED"
	//Modified is a modify event
	Modified EventType = "MODIFIED"
	//Deleted is a delete event
	Deleted EventType = "DELETED"
	//Error is a error event
	Error EventType = "ERROR"
)

// Event represents a single event to a watched resource.
type Event struct {
	Type   EventType `json:"type"`
	Object Endpoints `json:"object"`
}

//Endpoints represents information of a resource
type Endpoints struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Subsets    []Subset `json:"subsets"`
}

//Metadata represents some additional info of resource
type Metadata struct {
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	ResourceVersion string            `json:"resourceVersion"`
	Labels          map[string]string `json:"labels"`
}

//Subset represents addresses to a resource
type Subset struct {
	Addresses []Address `json:"addresses"`
	Ports     []Port    `json:"ports"`
}

//Address is a access address to a target
type Address struct {
	IP        string           `json:"ip"`
	TargetRef *ObjectReference `json:"targetRef,omitempty"`
}

//ObjectReference represents information of a target
type ObjectReference struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

//Port port info
type Port struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}
