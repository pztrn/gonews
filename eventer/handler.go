package eventer

// EventHandler represents generic event handler.
type EventHandler struct {
	// Command is a command name we will use for event registering,
	// showing help, etc.
	Command string
	// Handler if a function that will handle this command.
	Handler func(data interface{}) interface{}
}
