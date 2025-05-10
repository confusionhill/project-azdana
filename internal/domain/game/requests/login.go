package requests

import (
	"context"
	"net"
	"sync"

	"com.github/confusionhill-aqw-ps/internal/config"
	"com.github/confusionhill-aqw-ps/internal/domain/auth"
	control "com.github/confusionhill-aqw-ps/internal/domain/game/control"
	dto "com.github/confusionhill-aqw-ps/internal/model/dto/auth"
	"com.github/confusionhill-aqw-ps/internal/utilities"
)

type Login struct {
	cfg   *config.Config
	mutex *sync.Mutex
	world *control.World
	repo  *auth.Repository
}

func NewLogin(cfg *config.Config, mutex *sync.Mutex, world *control.World, repo *auth.Repository) *Login {
	return &Login{
		cfg:   cfg,
		mutex: mutex,
		world: world,
		repo:  repo,
	}
}

func (l *Login) Handle(conn net.Conn, packet dto.GameLoginRequestDTO) utilities.Packet {
	sendPacket := utilities.NewPacket()
	sendPacket.AddString("%xt%loginResponse%-1%")
	user, err := l.repo.LoginUser(context.Background(), dto.LoginUserRequestDTO{
		Username: packet.Body.Login.Nick,
		Password: packet.Body.Login.Pword,
	})
	if err != nil {
		sendPacket.AddString("0%-1%%User Data could not be retrieved. Please contact the Adventure Quest Worlds staff to resolve the issue.%")
		return sendPacket
	}
	l.world.AddConn(l.mutex, conn, *user)
	// `%xt%loginResponse%-1%true%1%$dukun%Welcome to AsyncQuest! An educational AQWorlds Private Server made with NodeJS.%2020-08-05T21:42:25%sNews=news.swft,sMap=map.swft,sBook=book.swft,sFBC=dldld,sAssets=asset.swf,sWTSandbox=false,gMenu=menu.swf%\0`
	serverMsg := "Welcome to Adventure Quest Worlds! An educational AQWorlds Private Server made with Go."
	msg := "true%-1%" + "%" + user.Username + "%" + serverMsg + "%" + user.Password //+ "%" + ""
	sendPacket.AddString(msg)
	return sendPacket
}
