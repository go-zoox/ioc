package ioc

var (
	// global is the global container
	global Container
)

func getGlobal() Container {
	if global == nil {
		global = New()
	}

	return global
}

// Register service by id
func Register(id string, service any) {
	getGlobal().Register(id, service)
}

// Get gets a service by id
func Get(id string) (any, bool) {
	return getGlobal().Get(id)
}

// MustGet calls Get underneath
// will panic if serviceect not found within container
func MustGet(id string) any {
	return getGlobal().MustGet(id)
}

// Invoke gets a service safely typed by passing it to a closure
// will panic if callback is not a function
func Invoke(id string, fn any) {
	getGlobal().Invoke(id, fn)
}

// MustInvoke calls MustGet underneath
// will panic if service not found within container
func MustInvoke(id string, fn any) {
	getGlobal().MustInvoke(id, fn)
}

// Has checks if a service exists within the container
func Has(id string) bool {
	return getGlobal().Has(id)
}
