package server

// Server is a simple micro server abstraction
type Server interface {
	Options() Options
	Init(...Option) error
}
