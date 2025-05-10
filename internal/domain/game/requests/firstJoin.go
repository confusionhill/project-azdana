package requests

import (
	"fmt"
	"net"
	"sync"

	"com.github/confusionhill-aqw-ps/internal/config"
	control "com.github/confusionhill-aqw-ps/internal/domain/game/control"
	"com.github/confusionhill-aqw-ps/internal/model/entity/game"
	"com.github/confusionhill-aqw-ps/internal/utilities"
)

type FirstJoin struct {
	cfg   *config.Config
	mutex *sync.Mutex
	world *control.World
	room  *game.Room
}

func NewFirstJoin(cfg *config.Config, mutex *sync.Mutex, world *control.World) *FirstJoin {
	return &FirstJoin{cfg: cfg, mutex: mutex, world: world}
}

// TODO: Bug fix this, figure out what those values are
func (f *FirstJoin) Handle(conn net.Conn) utilities.Packet {
	sendPack := utilities.NewPacket()
	player, err := f.world.GetPlayer(f.mutex, conn)
	if err != nil {
	}
	room, err := f.world.GetMap(f.mutex, "battleon", 0, player)
	if err != nil {
	}
	f.room = room
	sendPack.AddXMLSingle(1, "msg t", "sys")
	sendPack.AddXMLSingle(1, "body action", "joinOK", "r", "4")
	sendPack.AddXMLSingle(0, "pid id", "1") //fmt.Sprintf("%d", player.ID))
	sendPack.AddXMLSingle(0, "vars")
	sendPack.AddXMLSingle(1, "uLs r", "4")
	for idx, _ := range room.Players {
		mod := "1"
		strIdx := fmt.Sprintf("%d", idx+1)
		// strPlayerId := fmt.Sprintf("%d", roomPlayer.ID)
		sendPack.AddXMLSingle(1, "u i", "1", "m", mod, "s", "0", "p", strIdx)
		sendPack.AddXML("n", "", 1)
		sendPack.AddCDATA(player.Username)
		sendPack.AddXML("n", "", 2)
		sendPack.AddXML("vars", "", 0)
		sendPack.AddXMLSingle(2, "u")
	}
	sendPack.AddXMLSingle(2, "uLs")
	sendPack.AddXMLSingle(2, "body")
	sendPack.AddXMLSingle(2, "msg")
	return sendPack
}

func (f *FirstJoin) GetMap(conn net.Conn) utilities.Packet {
	sendPack := utilities.NewPacket()
	// player, err := f.world.GetPlayer(f.mutex, conn)
	// if err != nil {
	// }
	room := f.room
	sendPack.AddString("{\"t\":\"xt\",\"b\":{\"r\":-1,\"o\":{\"cmd\":\"moveToArea\",\"areaName\":\"")
	roomId := fmt.Sprintf("%d", room.Id)
	// areaId := fmt.Sprintf("%d", room.Map.Id)
	sendPack.AddString(room.Map.Name + "-" + roomId)
	sendPack.AddString("\",\"intKillCount\":0,\"uoBranch\":[")
	for idx, roomPlayer := range room.Players {
		// playerId := fmt.Sprintf("%d", roomPlayer.ID)
		if idx != 0 {
			sendPack.AddString(",")
		}
		sendPack.AddString(f.getPlayerInfo(roomPlayer))
	}
	sendPack.AddString("],\"strMapFileName\":\"")
	sendPack.AddString(room.Map.FileName)
	sendPack.AddString("\",\"intType\":\"2\",\"monBranch\":[],\"wB\":[],\"sExtra\":\"a1=test1,a2=test2\",\"areaId\":")
	//TODO: Bug fix this, figure out what those values are
	sendPack.AddString("4")
	sendPack.AddString(",\"strMapName\":\"")
	sendPack.AddString(room.Map.Name)
	sendPack.AddString("\"}}}")
	return sendPack
}

func (f *FirstJoin) getPlayerInfo(player game.User) string {
	pInfo := utilities.NewPacket()
	level := fmt.Sprintf("%d", player.Level)
	pInfo.AddString("[\"uoName:" + player.Username)
	pInfo.AddString("\",\"strUsername:" + player.Username)
	pInfo.AddString("\",\"strFrame:" + "Enter")
	pInfo.AddString("\",\"strPad:" + "Spawn")
	pInfo.AddString("\",\"intState:1\",\"intLevel:" + level)
	pInfo.AddString("\",\"intHP:" + "760")
	pInfo.AddString("\",\"intMP:" + "1000")
	pInfo.AddString("\",\"intHPMax:" + "10000")
	pInfo.AddString("\",\"intMPMax:" + "10000")
	pInfo.AddString("\",\"tx:" + "0")
	pInfo.AddString("\",\"ty:" + "0")
	pInfo.AddString("\",\"afk:" + "false")
	pInfo.AddString("\",\"vip:" + "0")
	pInfo.AddString("\",\"fd:" + "-1" + "\"]")
	return pInfo.GetPacket()
}

// func ()
