package handler

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"

	"com.github/confusionhill-aqw-ps/application/consumer"
	"com.github/confusionhill-aqw-ps/internal/config"
	"com.github/confusionhill-aqw-ps/internal/domain/game/control"
	"com.github/confusionhill-aqw-ps/internal/domain/game/requests"
	"com.github/confusionhill-aqw-ps/internal/model/dto/auth"
	"com.github/confusionhill-aqw-ps/internal/model/entity/game"
	"com.github/confusionhill-aqw-ps/internal/utilities"
	"github.com/labstack/gommon/log"
)

type Handlers struct {
	cfg          *config.Config
	conn         net.Conn
	mutex        *sync.Mutex
	reader       *bufio.Reader
	writer       *bufio.Writer
	world        *control.World
	loginRequest *requests.Login
}

func New(cfg *config.Config, conn net.Conn, mutex *sync.Mutex, world *control.World, repo *consumer.Repositories) Handlers {
	return Handlers{
		cfg:          cfg,
		conn:         conn,
		mutex:        mutex,
		reader:       bufio.NewReader(conn),
		writer:       bufio.NewWriter(conn),
		world:        world,
		loginRequest: requests.NewLogin(cfg, mutex, world, repo.Auth),
	}
}

func (h *Handlers) MainConnHandler() {
	defer func() {
		h.conn.Close()
		log.Debug("conn closed")
	}()
	h.world.AddConn(h.mutex, h.conn, game.User{})

	for {
		message, err := h.reader.ReadString('\x00')
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			break
		}
		packet := strings.TrimSuffix(message, "\x00")
		reqPacket := utilities.NewPacket()
		reqPacket.SetPacket(packet)
		fmt.Printf("from user[%d]: %s\n", 1, packet)
		cmd, _ := h.getCmd(packet)
		switch cmd {
		case CMD_POLICY_FILE_REQUEST:
			sendPack := utilities.NewPacket()
			sendPack.AddXML("cross-domain-policy", "", 1)
			sendPack.AddXMLSingle(0, "allow-access-from domain", "*", "to-ports", "5588")
			sendPack.AddXML("cross-domain-policy", "", 2)
			h.send(sendPack, true)
		case CMD_MSG_SYS_REQUEST:
			sendPack := utilities.NewPacket()
			if strings.Contains(packet, "verChk") {
				sendPack.AddXMLSingle(1, "msg t", "sys")
				sendPack.AddXMLSingle(1, "body action", "apiOK")
				sendPack.AddXMLSingle(2, "body")
				sendPack.AddXMLSingle(2, "msg")
				h.send(sendPack, true)
			} else if strings.Contains(packet, "login") {
				msg := auth.GameLoginRequestDTO{}
				err := xml.Unmarshal([]byte(packet), &msg)
				if err != nil {
					log.Error("error unmarshalling packet: %s", err.Error())
					break
				}
				sendPack := h.loginRequest.Handle(h.conn, msg)
				h.send(sendPack, true)
			}
		}
		fmt.Printf("cmd: %s\n", cmd)
	}
	fmt.Println("conn closed")
	h.world.RemoveConn(h.mutex, h.conn)
}

func (h *Handlers) getCmd(packet string) (string, error) {
	if strings.HasPrefix(packet, "<") {
		endArrow := strings.Index(packet, ">")
		endSlash := strings.Index(packet, "/>")
		if endSlash < 0 {
			endSlash = endArrow + 1
		}
		if endSlash < endArrow {
			return packet[1:endSlash], nil
		} else {
			return packet[1:endArrow], nil
		}
	} else if strings.HasPrefix(packet, "%") {
		packetHandled := strings.Split(packet, "%")
		if len(packetHandled) > 3 {
			return packetHandled[3], nil
		}
	}
	return "", errors.New("cannot get cmd")
}

func (h *Handlers) send(sendPack utilities.Packet, addNull bool) {
	packet := sendPack.GetPacket()
	if addNull {
		packet += "\u0000"
	}
	h.writer.Write([]byte(packet))
	h.writer.Flush()
}
