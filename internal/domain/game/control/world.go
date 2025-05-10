package control

import (
	"errors"
	"net"
	"sync"

	"com.github/confusionhill-aqw-ps/internal/model/entity/game"
	"com.github/confusionhill-aqw-ps/internal/model/errorsEntity"
	"github.com/labstack/gommon/log"
)

type NetPool map[net.Conn]game.User

const (
	MAX_CONN = 2000
)

type World struct {
	maps    map[string]map[int64]game.Room
	netPool NetPool
}

func NewWorld(maps []game.Map) World {
	resMaps := make(map[string]map[int64]game.Room)
	for _, locMap := range maps {
		room := map[int64]game.Room{
			1: {
				Id:      1,
				Map:     locMap,
				Players: make([]game.User, 0),
			},
		}
		resMaps[locMap.Name] = room
	}
	return World{
		maps:    resMaps,
		netPool: make(NetPool, 0),
	}
}

// selectedMap := w.Maps[name]
//
//	if roomId == 0 {
//		for _, room := range selectedMap {
//			if len(room.Players) < int(room.Map.MaxPlayer) {
//				return w.getPrivateRoom(name, roomId, player)
//			}
//		}
//	} else {just
//
//		return w.getPrivateRoom(name, roomId, player)
//	}
func (w *World) GetMap(mutex *sync.Mutex, name string, roomId int64, player game.User) (*game.Room, error) {
	return w.getPrivateRoom(mutex, name, 1, player)
}

func (w *World) getPrivateRoom(mutex *sync.Mutex, name string, roomId int64, player game.User) (*game.Room, error) {
	selectedMap := w.maps[name]
	currRoom := selectedMap[roomId]
	if playerCount := len(currRoom.Players); playerCount >= int(currRoom.Map.MaxPlayer) {
		return nil, errors.New(errorsEntity.ROOM_FULL)
	}
	for _, pl := range currRoom.Players {
		if pl.ID == player.ID {
			return nil, errors.New(errorsEntity.ROOM_PLAYER_INSIDE)
		}
	}
	newMap := currRoom
	newMap.Players = append(newMap.Players, player)
	w.maps[name][roomId] = newMap
	return &newMap, nil
}

func (w *World) IsMaxConn(mutex *sync.Mutex) bool {
	mutex.Lock()
	defer mutex.Unlock()
	return len(w.netPool) >= MAX_CONN
}

func (w *World) AddConn(mutex *sync.Mutex, conn net.Conn, player game.User) {
	mutex.Lock()
	w.netPool[conn] = player
	mutex.Unlock()
	log.Debugf("Added connection: %v", conn)
}

func (w *World) RemoveConn(mutex *sync.Mutex, conn net.Conn) {
	mutex.Lock()
	delete(w.netPool, conn)
	mutex.Unlock()
	log.Debugf("Removed connection: %v", conn)
}
