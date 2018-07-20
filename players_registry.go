package typhoon

import "sync"

var (
	playersCount                 = 0
	players      map[int]*Player = make(map[int]*Player)
	playersMutex                 = &sync.Mutex{}
)

func (player *Player) register() {
	playersMutex.Lock()
	players[player.id] = player
	playersCount++
	playersMutex.Unlock()
}

func (player *Player) unregister() {
	playersMutex.Lock()
	if _, ok := players[player.id]; ok {
		playersCount--
	}
	delete(players, player.id)
	playersMutex.Unlock()
}
