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
	defer reg.playersMutex.Unlock()
	reg.players[player.id] = player
	reg.playersCount++
}

func (player *Player) unregister() {
	reg := player.core.playerRegistry
	reg.playersMutex.Lock()
	defer reg.playersMutex.Unlock()
	if _, ok := reg.players[player.id]; ok {
		reg.playersCount--
	}
	delete(reg.players, player.id)
}

func (registry *PlayerRegistry) ForEachPlayer(fn func(player *Player)) {
	registry.playersMutex.RLock()
	defer registry.playersMutex.RUnlock()
	for _, player := range registry.players {
		fn(player)
	}
}

func (registry *PlayerRegistry) ForEachPlayerAsync(fn func(player *Player)) {
	registry.playersMutex.RLock()
	var wg sync.WaitGroup
	for _, player := range registry.players {
		wg.Add(1)
		go func(p *Player) {
			defer wg.Done()
			fn(p)
		}(player)
	}
	registry.playersMutex.RUnlock()
	wg.Wait()
}

func (registry *PlayerRegistry) GetPlayerCount() int {
	registry.playersMutex.RLock()
	defer registry.playersMutex.RUnlock()
	return registry.playersCount
}

func (registry *PlayerRegistry) GetPlayers() []*Player {
	registry.playersMutex.RLock()
	defer registry.playersMutex.RUnlock()
	players := make([]*Player, 0, len(registry.players))
	for _, pl := range registry.players {
		players = append(players, pl)
	}
	return players
}
