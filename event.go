package typhoon

import (
	"reflect"
)

type Event interface {
	GetID() string
}

const (
	TyphoonEventPlayerJoin = "typhoon:PlayerJoinEvent"
	TyphoonEventPlayerChat = "typhoon:PlayerChatEvent"
)

type PlayerJoinEvent struct {
	Player *Player
}

func (e *PlayerJoinEvent) GetID() string {
	return TyphoonEventPlayerJoin
}

type PlayerChatEvent struct {
	Player  *Player
	Message string
}

func (e *PlayerChatEvent) GetID() string {
	return TyphoonEventPlayerChat
}

func (c *Core) On(handler interface{}) {
	typ := reflect.TypeOf(handler).In(0)
	ev := reflect.New(typ).Elem().Interface().(Event)

	if arr, ok := c.eventHandlers[ev.GetID()]; ok {
		c.eventHandlers[ev.GetID()] = append(arr, handler)
	} else {
		arr := make([]interface{}, 1)
		arr[0] = handler
		c.eventHandlers[ev.GetID()] = arr
	}
}

func (c *Core) CallEvent(event Event) {
	for _, f := range c.eventHandlers[event.GetID()] {
		reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(event)})
	}
}
