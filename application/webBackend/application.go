package webbackend

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"com.github/confusionhill-aqw-ps/application/consumer"
	"com.github/confusionhill-aqw-ps/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserForm struct {
	Email     string `form:"strEmail"`
	Age       int    `form:"intAge"`
	Gender    string `form:"strGender"`
	HairColor int    `form:"intColorHair"`
	SkinColor int    `form:"intColorSkin"`
	HairID    int    `form:"HairID"`
	EyeColor  int    `form:"intColorEye"`
	DOB       string `form:"strDOB"`
	Password  string `form:"strPassword"`
	Username  string `form:"strUsername"`
	ClassID   int    `form:"ClassID"`
}

func generateSessionId() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		time.Now().UnixNano()&0xFFFFFFFF,
		time.Now().UnixNano()>>32&0xFFFF,
		(time.Now().UnixNano()>>48&0x0FFF)|0x4000, // Ensure RFC4122 version 4 UUID
		(time.Now().UnixNano()>>60&0x3FFF)|0x8000, // Ensure RFC4122 variant
		time.Now().UnixNano()>>64&0xFFFFFFFFFFFF,
	)
}

const (
	handshakeToken = "#"
	disconnectCmd  = "disconnect"
	connectCmd     = "connect"
	paramName      = "sfsHttp"
	servletPath    = "/BlueBox/HttpBox.do"
	connectionLost = "ERR#01"
)

var sessions = make(map[string]time.Time) // sessionId -> lastActiveTime

func RunWebBackendApp(cfg *config.Config, handlers *consumer.Handlers) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/gamefiles", "public/gamefiles")
	e.Static("/assets", "public/pages/assets")
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/game")
	})

	e.GET("/game", func(c echo.Context) error {
		return c.File("public/pages/game.html")
	})

	e.GET("/register", func(c echo.Context) error {
		return c.File("public/pages/register.html")
	})

	e.GET("/getversion.asp", func(c echo.Context) error {
		return c.String(http.StatusOK, "status=success&sFile=game_port_8080.swf&sTitle=Alpha Test&sBG=bg.swf")
	})

	e.POST("/game/cf-userlogin.asp", handlers.Auth.LoginUserHandler)
	e.POST("/game/cf-userlogin.php", handlers.Auth.LoginUserHandler)
	e.POST("/cf-userlogin.php", handlers.Auth.LoginUserHandler)
	e.POST("/cf-usersignup.php", handlers.Auth.RegisterUserHandler)
	e.POST("/BlueBox/HttpBox.do", func(c echo.Context) error {
		sfsData := c.FormValue(paramName)
		if sfsData == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing 'sfsHttp' parameter")
		}

		parts := strings.SplitN(sfsData, "|", 2)
		var sessionId string
		var command string

		if len(parts) == 2 {
			sessionId = parts[0]
			command = parts[1]
		} else if len(parts) == 1 {
			command = parts[0]
		}

		fmt.Printf("Received command: '%s' with sessionId: '%s'\n", command, sessionId)

		switch command {
		case connectCmd:
			newSessionId := generateSessionId()
			sessions[newSessionId] = time.Now()
			response := handshakeToken + newSessionId
			fmt.Printf("Sent handshake: %s\n", response)
			return c.String(http.StatusOK, response)
		case disconnectCmd:
			if _, ok := sessions[sessionId]; ok {
				delete(sessions, sessionId)
				fmt.Printf("Session %s disconnected.\n", sessionId)
				return c.String(http.StatusOK, connectionLost)
			} else {
				fmt.Printf("Attempted disconnect with invalid session: %s\n", sessionId)
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid session")
			}
		default:
			if _, ok := sessions[sessionId]; ok {
				sessions[sessionId] = time.Now() // Update last active time
				fmt.Printf("Echoed back for session %s: %s\n", sessionId, command)
				return c.String(http.StatusOK, command)
			} else {
				fmt.Printf("Received data with invalid session: %s\n", sessionId)
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid session")
			}
		}
	})

	return e.Start(cfg.Server.WebPort)
}
