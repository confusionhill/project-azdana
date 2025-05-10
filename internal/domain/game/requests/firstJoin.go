package requests

import (
	"sync"

	control "com.github/confusionhill-aqw-ps/internal/domain/game/control"
	"com.github/confusionhill-aqw-ps/internal/model/entity/game"
)

type FirstJoin struct {
	world *control.World
}

func (f *FirstJoin) Request(mutex *sync.Mutex, player game.User) {
	f.world.GetMap(mutex, "battleon", 0, player)
}
