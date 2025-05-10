package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"

	"com.github/confusionhill-aqw-ps/application"
	gameemulator "com.github/confusionhill-aqw-ps/application/gameEmulator"
	"com.github/confusionhill-aqw-ps/internal/config"
)

const (
	SERVER_PORT = "5588"
	NEWS        = "Some news"
	MAP         = "Battleon"
	BOOK        = "Book of Lore"
	FBC         = "false"
	ASSETS      = "http://localhost/assets/"
	GMENU       = "defaultMenu"
)

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
}

var (
	users       = make([]map[string]interface{}, 2)
	onlineCount = 0
	sessId      = 2
	mu          sync.Mutex
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	rsc, repo, _, _, err := application.RunApplication(cfg)
	if err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
	err = gameemulator.RunGameEmulator(cfg, rsc, repo)
	if err != nil {
		log.Fatalf("Failed to run game emulator: %v", err)
	}
}

func tcp() {
	ln, err := net.Listen("tcp", ":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println(`
|----------------------|
|                      |
|                      |
|        GoLand	       |
|         v0.1         |
|                      |
|----------------------|`)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		buffer, err := reader.ReadString('\x00')
		if err == io.EOF {
			fmt.Println("disconnected.")
			return
		}
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		packet := strings.TrimSuffix(buffer, "\x00")
		fmt.Printf("from user[%d]: %s\n", sessId, packet)
		conn.Write([]byte("Hello, client!"))
	}
}
