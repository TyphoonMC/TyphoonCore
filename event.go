package typhoon

import (
	"reflect"
)

type EventCallback struct {
	Callback interface{}
	MetaData map[string]string
}
type Event interface{}

// Cancellable interface for events that can be cancelled
type Cancellable interface {
	IsCancelled() bool
	SetCancelled(cancelled bool)
}

// BaseCancellable provides default implementation for cancellable events
type BaseCancellable struct {
	cancelled bool
}

func (b *BaseCancellable) IsCancelled() bool {
	return b.cancelled
}

func (b *BaseCancellable) SetCancelled(cancelled bool) {
	b.cancelled = cancelled
}

// ============================================
// Connection Events
// ============================================

// PlayerPreLoginEvent is called when a player attempts to connect (async, before login)
type PlayerPreLoginEvent struct {
	BaseCancellable
	Player     *Player
	Name       string
	Address    string
	KickReason string
}

// PlayerLoginEvent is called when a player logs in
type PlayerLoginEvent struct {
	BaseCancellable
	Player     *Player
	KickReason string
}

// PlayerPreJoinEvent is called before a player joins the server
type PlayerPreJoinEvent struct {
	Player *Player
}

// PlayerJoinEvent is called when a player joins the server
type PlayerJoinEvent struct {
	Player      *Player
	JoinMessage string
}

// PlayerQuitEvent is called when a player leaves the server
type PlayerQuitEvent struct {
	Player      *Player
	QuitMessage string
}

// PlayerKickEvent is called when a player is kicked from the server
type PlayerKickEvent struct {
	BaseCancellable
	Player *Player
	Reason string
}

// ============================================
// Chat & Command Events
// ============================================

// PlayerChatEvent is called when a player sends a chat message
type PlayerChatEvent struct {
	BaseCancellable
	Player  *Player
	Message string
	Format  string
}

// PlayerCommandPreprocessEvent is called when a player executes a command
type PlayerCommandPreprocessEvent struct {
	BaseCancellable
	Player  *Player
	Command string
}

// PlayerTabCompleteEvent is called when a player uses tab completion
type PlayerTabCompleteEvent struct {
	Player      *Player
	Buffer      string
	Completions []string
}

// ============================================
// Movement Events
// ============================================

// PlayerMoveEvent is called when a player moves
type PlayerMoveEvent struct {
	BaseCancellable
	Player *Player
	From   Location
	To     Location
}

// PlayerTeleportEvent is called when a player teleports
type PlayerTeleportEvent struct {
	BaseCancellable
	Player *Player
	From   Location
	To     Location
	Cause  TeleportCause
}

// TeleportCause represents the reason for a teleport
type TeleportCause int

const (
	TeleportCauseUnknown TeleportCause = iota
	TeleportCauseCommand
	TeleportCausePlugin
	TeleportCauseEnderPearl
	TeleportCauseNetherPortal
	TeleportCauseEndPortal
	TeleportCauseSpectate
)

// ============================================
// Interaction Events
// ============================================

// PlayerClickType represents the type of click
type PlayerClickType byte

const (
	PlayerRightClick PlayerClickType = iota
	PlayerLeftClick
)

// PlayerInteractEvent is called when a player interacts
type PlayerInteractEvent struct {
	BaseCancellable
	Player    *Player
	ClickType PlayerClickType
	BlockX    int
	BlockY    int
	BlockZ    int
}

// PlayerAnimationEvent is called when a player performs an animation
type PlayerAnimationEvent struct {
	BaseCancellable
	Player        *Player
	AnimationType AnimationType
}

// AnimationType represents player animations
type AnimationType int

const (
	AnimationArmSwing AnimationType = iota
	AnimationCriticalEffect
	AnimationMagicCriticalEffect
	AnimationLeaveBed
)

// ============================================
// State Change Events
// ============================================

// PlayerGameModeChangeEvent is called when a player's gamemode changes
type PlayerGameModeChangeEvent struct {
	BaseCancellable
	Player      *Player
	NewGameMode Gamemode
}

// PlayerToggleFlightEvent is called when a player toggles flight
type PlayerToggleFlightEvent struct {
	BaseCancellable
	Player    *Player
	IsFlying  bool
}

// PlayerToggleSneakEvent is called when a player toggles sneak
type PlayerToggleSneakEvent struct {
	BaseCancellable
	Player     *Player
	IsSneaking bool
}

// PlayerToggleSprintEvent is called when a player toggles sprint
type PlayerToggleSprintEvent struct {
	BaseCancellable
	Player      *Player
	IsSprinting bool
}

// ============================================
// Inventory Events
// ============================================

// PlayerDropItemEvent is called when a player drops an item
type PlayerDropItemEvent struct {
	BaseCancellable
	Player *Player
	// ItemStack ItemStack // TODO: implement ItemStack
}

// PlayerHeldItemChangeEvent is called when a player changes held item slot
type PlayerHeldItemChangeEvent struct {
	BaseCancellable
	Player       *Player
	PreviousSlot int
	NewSlot      int
}

// ============================================
// World Events
// ============================================

// PlayerChangedWorldEvent is called when a player changes world
type PlayerChangedWorldEvent struct {
	Player *Player
	From   string
	To     string
}

// PlayerRespawnEvent is called when a player respawns
type PlayerRespawnEvent struct {
	Player           *Player
	RespawnLocation  Location
	IsBedSpawn       bool
	IsAnchorSpawn    bool
}

// ============================================
// Resource Pack Events
// ============================================

// PlayerResourcePackStatusEvent is called when a player responds to resource pack request
type PlayerResourcePackStatusEvent struct {
	Player *Player
	Status ResourcePackStatus
}

// ResourcePackStatus represents the status of a resource pack request
type ResourcePackStatus int

const (
	ResourcePackSuccessfullyLoaded ResourcePackStatus = iota
	ResourcePackDeclined
	ResourcePackFailedDownload
	ResourcePackAccepted
)

// ============================================
// Plugin Message Events
// ============================================

// PluginMessageEvent is called when a plugin message is received
type PluginMessageEvent struct {
	Player  *Player
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
