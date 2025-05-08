package webbackend

import (
	"net/http"

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
		return c.String(http.StatusOK, "status=success&sFile=client/game_ori_rev.swf&sTitle=Alpha Test&sBG=bg.swf")
	})

	e.POST("/game/cf-userlogin.asp", handlers.Auth.LoginUserHandler)
	e.POST("/cf-usersignup.php", handlers.Auth.RegisterUserHandler)

	return e.Start(cfg.Server.WebPort)
}
