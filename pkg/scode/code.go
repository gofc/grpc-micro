package scode

//ServiceCode an unique service code
type ServiceCode string

const (
	//FOO test service(FooService)
	FOO ServiceCode = "foo"
)

//Name get service code
func (s ServiceCode) Name() string {
	return string(s)
}
