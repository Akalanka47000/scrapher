package global

var shutdownHooks map[string]func()

// Register a function to be called on server shutdown.
// This differs from the fiber shutdown hook since these will be called some time after
// the server has stopped accepting new requests and when it's time to kill the process after gracefully
// allowing all requests to finish.
func RegisterShutdownHook(name string, hook func()) {
	if shutdownHooks == nil {
		shutdownHooks = make(map[string]func())
	}
	shutdownHooks[name] = hook
}

// Unregister a function from being called on server shutdown.
func UnregisterShutdownHook(name string) {
	if shutdownHooks == nil {
		return
	}
	delete(shutdownHooks, name)
}

// Execute all registered shutdown hooks. This is called just before the process exits.
func ExecuteShutdownHooks() {
	if shutdownHooks == nil {
		return
	}
	for _, hook := range shutdownHooks {
		hook()
	}
}
