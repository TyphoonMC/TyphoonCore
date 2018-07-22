package typhoon

import (
	"reflect"
)

type Event interface {}

type PlayerJoinEvent struct {
	Player *Player
}

type PlayerChatEvent struct {
	Player  *Player
	Message string
}

func (c *Core) On(handler interface{}) {
	typ := reflect.TypeOf(handler).In(0)

	if arr, ok := c.eventHandlers[typ]; ok {
		c.eventHandlers[typ] = append(arr, handler)
	} else {
		arr := make([]interface{}, 1)
		arr[0] = handler
		c.eventHandlers[typ] = arr
	}
}

func (c *Core) CallEvent(event Event) {
	for _, f := range c.eventHandlers[reflect.TypeOf(event)] {
		reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(event)})
	}
}
