package ioc

// New creates a new DI Container
func New() Container {
	return &container{
		registry: make(map[string]any),
	}
}
