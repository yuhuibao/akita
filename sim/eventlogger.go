package sim

import (
	"log"
	"reflect"
)

// EventLogger is an hook that prints the event information
type EventLogger struct {
	LogHookBase
}

// NewEventLogger returns a new LogEventHook which will write in to the logger
func NewEventLogger(logger *log.Logger) *EventLogger {
	h := new(EventLogger)
	h.Logger = logger
	return h
}

// Func writes the event information into the logger
func (h *EventLogger) Func(ctx HookCtx) {
	if ctx.Pos != HookPosBeforeEvent {
		return
	}

	evt, ok := ctx.Item.(Event)
	if !ok {
		return
	}
	time, ok := ctx.Detail.(int64)
	if !ok {
		return
	}

	comp, ok := evt.Handler().(Component)
	if ok {
		h.Logger.Printf("%.10f, %s -> %s %d",
			evt.Time(), reflect.TypeOf(evt), comp.Name(), time)
	} else {
		h.Logger.Printf("%.10f, %s", evt.Time(), reflect.TypeOf(evt))
	}
}
