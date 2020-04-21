package registry

//Config of registry
type Config struct {
	Type    string `yaml:"type" default:"etcd"`
	Address string `yaml:"address" required:"true"`
}
