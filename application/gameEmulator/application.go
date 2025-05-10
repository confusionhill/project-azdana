package gameemulator

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"com.github/confusionhill-aqw-ps/application/consumer"
	"com.github/confusionhill-aqw-ps/internal/config"
	control "com.github/confusionhill-aqw-ps/internal/domain/game/control"
	"com.github/confusionhill-aqw-ps/internal/domain/game/handler"
	"com.github/confusionhill-aqw-ps/internal/model/entity/game"
	"github.com/labstack/gommon/log"
)

func RunGameEmulator(cfg *config.Config, rsc *consumer.Resources, repo *consumer.Repositories) error {
	world := control.NewWorld(make([]game.Map, 0))
	net, err := net.Listen("tcp", cfg.Server.GamePort)
	if err != nil {
		return err
	}
	defer net.Close()
	return startServer(cfg, net, &world, repo)
}

func startServer(cfg *config.Config, network net.Listener, world *control.World, repo *consumer.Repositories) error {
	var mutex sync.Mutex
	log.Info(fmt.Sprintf("start server: %s", cfg.Server.GamePort))
	for {
		conn, err := network.Accept()
		if err != nil {
			log.Error(err)
			continue
		}
		if world.IsMaxConn(&mutex) {
			conn.Close()
			continue
		}
		hdl := handler.New(cfg, conn, &mutex, world, repo)
		go hdl.MainConnHandler()
	}
	return errors.New("error")
}
