package requests

import (
	"context"
	"net"
	"strings"
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

type loginResponse struct {
	Username string
	Message  string
	Password string
}

func (lr loginResponse) toMessage() string {
	template := "%xt%loginResponse%-1%true%1%{username}%{message}%{password}%sNews=news/News-Aug14.swf,sMap=news/Map-Aug14.swf,sBook=news/Book-July16.swf%"

	// Replace placeholders
	result := strings.ReplaceAll(template, "{username}", lr.Username)
	result = strings.ReplaceAll(result, "{message}", lr.Message)
	result = strings.ReplaceAll(result, "{password}", lr.Password)

	return result
}

func (l *Login) Handle(conn net.Conn, packet dto.GameLoginRequestDTO) utilities.Packet {
	sendPacket := utilities.NewPacket()
	// sendPacket.AddString("%xt%loginResponse%-1%")
	user, err := l.repo.LoginUser(context.Background(), dto.LoginUserRequestDTO{
		Username: packet.Body.Login.Nick,
		Password: packet.Body.Login.Pword,
	})
	if err != nil {
		sendPacket.AddString("0%-1%%User Data could not be retrieved. Please contact the Adventure Quest Worlds staff to resolve the issue.%")
		return sendPacket
	}
	l.world.AddConn(l.mutex, conn, *user)
	loginResponse := loginResponse{
		Username: user.Username,
		Message:  "Welcome to Adventure Quest Worlds! An educational AQWorlds Private Server made with Go.",
		Password: user.Password,
	}
	msg := loginResponse.toMessage()
	sendPacket.AddString(msg)
	return sendPacket
}
