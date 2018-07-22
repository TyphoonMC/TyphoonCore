package typhoon

import (
	"reflect"
)

type EventCallback struct {
	Callback interface{}
	MetaData map[string]string
}
type Event interface{}

type PlayerJoinEvent struct {
	Player *Player
}

type PlayerChatEvent struct {
	Player  *Player
	Message string
}

type PluginMessageEvent struct {
	Channel string
	Data    []byte
}

func (c *Core) On(handler interface{}) {
	typ := reflect.TypeOf(handler).In(0)
	cb := EventCallback{handler, nil}

	if arr, ok := c.eventHandlers[typ]; ok {
		c.eventHandlers[typ] = append(arr, cb)
	} else {
		arr := make([]EventCallback, 1)
		arr[0] = cb
		c.eventHandlers[typ] = arr
	}
}

func (c *Core) OnPluginMessage(channel string, handler interface{}) {
	typ := reflect.TypeOf(handler).In(0)
	cb := EventCallback{handler, make(map[string]string)}
	cb.MetaData["channel"] = channel

	if arr, ok := c.eventHandlers[typ]; ok {
		c.eventHandlers[typ] = append(arr, cb)
	} else {
		arr := make([]EventCallback, 1)
		arr[0] = cb
		c.eventHandlers[typ] = arr
	}
}

func (c *Core) callEventInternal(callback interface{}, event Event) {
	reflect.ValueOf(callback).Call([]reflect.Value{reflect.ValueOf(event)})
}

func (c *Core) CallEvent(event Event) {
	typ := reflect.TypeOf(event)
	for _, f := range c.eventHandlers[typ] {
		if f.MetaData == nil {
			c.callEventInternal(f.Callback, event)
		} else {
			switch event.(type) {
			case *PluginMessageEvent:
				if f.MetaData["channel"] == event.(*PluginMessageEvent).Channel {
					c.callEventInternal(f.Callback, event)
				}
			}
		}
	}
}
