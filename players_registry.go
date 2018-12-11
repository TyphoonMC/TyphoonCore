package typhoon

import "sync"

type PlayerRegistry struct {
	playersCount int
	players      map[int]*Player
	playersMutex *sync.RWMutex
}

func newPlayerRegistry() *PlayerRegistry {
	return &PlayerRegistry{
		playersCount: 0,
		players:      make(map[int]*Player),
		playersMutex: &sync.RWMutex{},
	}
}

func (player *Player) register() {
	reg := player.core.playerRegistry
	reg.playersMutex.Lock()
	reg.players[player.id] = player
	reg.playersCount++
	reg.playersMutex.Unlock()
}

func (player *Player) unregister() {
	reg := player.core.playerRegistry
	reg.playersMutex.Lock()
	if _, ok := reg.players[player.id]; ok {
		reg.playersCount--
	}
	delete(reg.players, player.id)
	reg.playersMutex.Unlock()
}

func (registry *PlayerRegistry) ForEachPlayer(fn func(player *Player)) {
	registry.playersMutex.Lock()
	for _, player := range registry.players {
		fn(player)
	}
	registry.playersMutex.Unlock()
}

func (registry *PlayerRegistry) ForEachPlayerAsync(fn func(player *Player)) {
	for _, player := range registry.players {
		go fn(player)
	}
}

func (registry *PlayerRegistry) GetPlayerCount() int {
	registry.playersMutex.RLock()
	i := registry.playersCount
	registry.playersMutex.RUnlock()
	return i
}

func (registry *PlayerRegistry) GetPlayers() []*Player {
	registry.playersMutex.RLock()
	players := make([]*Player, len(registry.players))
	i := 0
	for _, pl := range registry.players {
		players[i] = pl
		i++
	}
	registry.playersMutex.RUnlock()
	return players
}
