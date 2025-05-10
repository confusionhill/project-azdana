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
	cfg              *config.Config
	conn             net.Conn
	mutex            *sync.Mutex
	reader           *bufio.Reader
	writer           *bufio.Writer
	world            *control.World
	loginRequest     *requests.Login
	firstJoinRequest *requests.FirstJoin
}

func New(cfg *config.Config, conn net.Conn, mutex *sync.Mutex, world *control.World, repo *consumer.Repositories) Handlers {
	return Handlers{
		cfg:              cfg,
		conn:             conn,
		mutex:            mutex,
		reader:           bufio.NewReader(conn),
		writer:           bufio.NewWriter(conn),
		world:            world,
		loginRequest:     requests.NewLogin(cfg, mutex, world, repo.Auth),
		firstJoinRequest: requests.NewFirstJoin(cfg, mutex, world),
	}
}

type FakeMsg struct {
	XMLName xml.Name `xml:"msg"`
	Text    string   `xml:",chardata"`
	T       string   `xml:"t,attr"`
	Body    struct {
		Text   string `xml:",chardata"`
		Action string `xml:"action,attr"`
		R      string `xml:"r,attr"`
		Login  struct {
			Text  string `xml:",chardata"`
			Z     string `xml:"z,attr"`
			Nick  string `xml:"nick"`
			Pword string `xml:"pword"`
		} `xml:"login"`
	} `xml:"body"`
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
				var req auth.GameLoginRequestDTO
				err := xml.Unmarshal([]byte(packet), &req)
				if err != nil {
					log.Error("error unmarshalling packet: %s", err.Error())
					break
				}
				sendPack := h.loginRequest.Handle(h.conn, req)
				h.send(sendPack, true)
			}
		case CMD_FIRST_JOIN:
			joinOkPack := utilities.NewPacket()
			jkw := h.firstJoinRequest.Handle(h.conn)
			// <msg t='sys'><body action='joinOK' r='1'><pid id='2' /><vars /><uLs r='2'><u i='2' m='0' s='0' p='1'><n><![CDATA[dukun]]></n><vars></vars></u></uLs></body></msg>
			// <msg t='sys'><body action='joinOK' r='4'><pid id='1' /><vars /><uLs r='4'><u i='1' m='0' s='0' p='1'><n><![CDATA[dukun]]></n><vars></vars></u></uLs></body></msg>
			joinOkPack.AddString("<msg t='sys'><body action='joinOK' r='4'><pid id='1' /><vars /><uLs r='4'><u i='1' m='0' s='0' p='1'><n><![CDATA[dukun]]></n><vars></vars></u></uLs></body></msg>\u0000")
			// joinOkPack.AddString("<msg t='sys'><body action='joinOK' r='1'><pid id='0' /><vars /><uLs r='0'><u i='0' m='0' s='0' p='1'><n><![CDATA[]]></n><vars></vars></u></uLs></body></msg>\u0000")
			fmt.Println("jkw1: ", jkw)
			fmt.Println("jkw2: ", joinOkPack)
			h.send(jkw, true)
			msg := "{\"t\":\"xt\",\"b\":{\"r\":-1,\"o\":{\"cmd\":\"moveToArea\",\"areaName\":\"battleon-1\",\"intKillCount\":0,\"uoBranch\":[[\"uoName:dukun\",\"strUsername:dukun\",\"strFrame:Enter\",\"strPad:Spawn\",\"intState:1\",\"intLevel:2\",\"intHP:760\",\"intMP:21\",\"intHPMax:760\",\"intMPMax:21\",\"tx:0\",\"ty:0\",\"afk:false\",\"vip:0\",\"fd:-1\"]],\"strMapFileName\":\"/Battleon/town-battleon-Aug14.swf\",\"intType\":\"2\",\"monBranch\":[],\"wB\":[],\"sExtra\":\"a1=test1,a2=test2\",\"areaId\":4,\"strMapName\":\"battleon\"}}}\u0000"
			sendMapPack := h.firstJoinRequest.GetMap(h.conn)
			fmt.Println("sendMapPack1: ", sendMapPack.GetPacket())
			fmt.Println("sendMapPack2: ", msg)
			h.send(sendMapPack, true)
			// h.writer.Write([]byte(msg))
			// h.writer.Flush()
			msgPack := utilities.NewPacket()
			msgPack.AddString("%xt%server%-1%You joined \"battleon-1\"!%\u0000")
			h.send(msgPack, true)
		case CMD_GET_USER_DATA:
			sendPack := utilities.NewPacket()
			msg := `{"t":"xt","b":{"r":-1,"o":{"uid":1,"strFrame":"Enter","cmd":"initUserData","strPad":"Spawn","data":{"intColorAccessory":"0","iCP":101,"intLevel":"2","iBagSlots":20,"ig0":0,"iUpgDays":"-0","intColorBase":"0","iSTR":"0","ip0":0,"iq0":0,"iAge":"55","iWIS":"0","intExpToLevel":"350","intGold":10214,"intMP":21,"iBankSlots":0,"iHouseSlots":20,"id0":0,"intColorSkin":"13088131","intMPMax":21,"intHPMax":760,"dUpgExp":"2009-01-20T17:53:00","iUpg":"0","CharID":"2","strClassName":"Warrior Class","iINT":"0","ItemID":"3","lastArea":"","intColorTrim":"0","intDBExp":0,"intExp":0,"UserID":"2","ia1":"0","ia0":0,"intHP":760,"strQuests":"000000000000000000000QT000000000000000000000000000","bitSuccess":"1","strHairName":"Default","intColorEye":"91294","iLCK":"0","eqp":{"ar":{"ItemID":"16","sFile":"warrior_skin.swf","sLink":"Warrior"},"Weapon":{"ItemID":"1","sFile":"items/swords/sword01.swf","sLink":"Sword01"},"co":{"ItemID":"774","sFile":"peasant2_skin.swf","sLink":"Peasant2"}},"iDBCP":101,"intDBGold":0,"intActivationFlag":"5","intAccessLevel":"5","strHairFilename":"hair/M/Default.swf","intColorHair":"6180663","HairID":"52","strGender":"M","strUsername":"dukun","iDEX":"0","intCoins":1000,"iEND":"0","strMapName":"battleon"}}}`
			sendPack.AddString(msg)
			h.send(sendPack, true)
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
	const nullRef = "\u0000"
	packet := sendPack.GetPacket()
	if addNull {
		packet += nullRef
	}
	h.writer.Write([]byte(packet))
	h.writer.Flush()
}
