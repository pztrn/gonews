package eventer

import (
	// stdlib
	"testing"

	// other
	"github.com/stretchr/testify/require"
)

func TestEventerInitializationAndShutdown(t *testing.T) {
	Initialize()
	require.NotNil(t, events)

	Shutdown()
	require.Nil(t, events)
	require.False(t, eventsInitialized)
}

func TestEventerAddEventHandler(t *testing.T) {
	Initialize()
	require.NotNil(t, events)

	handler := &EventHandler{
		Command: "TEST",
		Handler: func(data interface{}) interface{} {
			return nil
		},
	}
	AddEventHandler(handler)

	hndl, exists := events[handler.Command]
	require.Len(t, events, 1)
	require.True(t, exists)
	require.Equal(t, handler, hndl)

	Shutdown()
	require.Nil(t, events)
	require.False(t, eventsInitialized)
}

func TestEventerAddEventHandlerAfterInitializationCompleted(t *testing.T) {
	Initialize()
	require.NotNil(t, events)
	InitializeCompleted()

	handler := &EventHandler{
		Command: "TEST",
		Handler: func(data interface{}) interface{} {
			return nil
		},
	}
	AddEventHandler(handler)

	hndl, exists := events[handler.Command]
	require.Len(t, events, 0)
	require.False(t, exists)
	require.Nil(t, hndl)

	Shutdown()
	require.Nil(t, events)
	require.False(t, eventsInitialized)
}

func TestEventerLaunchExistingEvent(t *testing.T) {
	Initialize()
	require.NotNil(t, events)

	handler := &EventHandler{
		Command: "TEST",
		Handler: func(data interface{}) interface{} {
			return true
		},
	}
	AddEventHandler(handler)

	hndl, exists := events[handler.Command]
	require.Len(t, events, 1)
	require.True(t, exists)
	require.Equal(t, handler, hndl)

	data, err := LaunchEvent(handler.Command, nil)
	if err != nil {
		t.Fatal("Test event launch failed:", err.Error())
	}

	switch data.(type) {
	case bool:
		break
	default:
		t.Fatalf("Test event returned unacceptable data type: %T", data)
	}

	Shutdown()
	require.Nil(t, events)
	require.False(t, eventsInitialized)
}

func TestEventerLaunchNotExistingEvent(t *testing.T) {
	Initialize()
	require.NotNil(t, events)

	handler := &EventHandler{
		Command: "TEST",
		Handler: func(data interface{}) interface{} {
			return true
		},
	}
	AddEventHandler(handler)

	hndl, exists := events[handler.Command]
	require.Len(t, events, 1)
	require.True(t, exists)
	require.Equal(t, handler, hndl)

	data, err := LaunchEvent(handler.Command+"notexisting", nil)
	if err == nil {
		t.Fatal("LaunchEvent() returned empty error!")
	}

	require.Nil(t, data)

	Shutdown()
	require.Nil(t, events)
	require.False(t, eventsInitialized)
}
