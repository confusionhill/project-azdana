package webbackend

import (
	"net/http"

	"com.github/confusionhill-aqw-ps/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunWebBackendApp(cfg *config.Config) {
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
		return c.String(http.StatusOK, "status=success&sFile=client/game_ori_rev.swf&sTitle=Alpha Test&sBG=bg.swf")
	})

	e.POST("/game/cf-userlogin.asp", func(c echo.Context) error {
		return c.String(http.StatusOK, "<login bSuccess='0' sMsg='The username and password you entered did not match. Please check the spelling and try again.'/>")
	})

	e.Logger.Fatal(e.Start(cfg.Server.WebPort))
}
