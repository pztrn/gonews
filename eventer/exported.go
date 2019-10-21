package eventer

import (
	// stdlib
	"log"
)

var (
	// List of known event handler.
	// Format:
	//   - commands/CMDNAME - base handler for command. This what
	//     networker will execute on any new command. Case-insensitive.
	//   - internals/CMDNAME - all other internal things like database
	//     access.
	// This map will not be changed after gonewsd initialization is
	// complete.
	events map[string]*EventHandler
	// Flag that indicates that we have completed events mapping
	// initialization.
	// RWMutex can be here, but we also need to write to log if we have
	// completed initialization.
	eventsInitialized bool
)

// Initialize initializes package.
func Initialize() {
	log.Println("Initializing event handler...")

	events = make(map[string]*EventHandler)
}

// InitializeCompleted sets eventsInitialized bool to true to avoid
// all further events mapping changes.
func InitializeCompleted() {
	eventsInitialized = true

	log.Println("Events initialization completed")
}

// AddEventHandler adds event handler to a list of known handlers.
func AddEventHandler(event *EventHandler) {
	if !eventsInitialized {
		events[event.Command] = event
	} else {
		log.Println("Can't add event '" + event.Command + "' to a list of known events - events mapping initialization already completed, no more changes allowed.")
	}
}

// LaunchEvent launches desired event in synchronous manner.
func LaunchEvent(eventName string, data interface{}) (interface{}, error) {
	handler, exists := events[eventName]
	if !exists {
		return nil, ErrEventNotFound
	}

	returnedData := handler.Handler(data)

	return returnedData, nil
}

// Shutdown just nullifies eventer map.
// Useful for testing and, maybe, reload.
func Shutdown() {
	events = nil
	eventsInitialized = false
}
