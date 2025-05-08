package auth

import (
	"net/http"

	"com.github/confusionhill-aqw-ps/internal/config"
	"com.github/confusionhill-aqw-ps/internal/model/dto/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	cfg     *config.Config
	usecase *Usecase
}

func NewHandler(cfg *config.Config, usecase *Usecase) (*Handler, error) {
	return &Handler{cfg: cfg, usecase: usecase}, nil
}

func (h *Handler) LoginUserHandler(c echo.Context) error {
	var req auth.LoginUserRequestDTO
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return c.String(http.StatusOK, "<login bSuccess='0' sMsg='The username and password you entered did not match. Please check the spelling and try again.'/>")
	}
	user, err := h.usecase.loginUserUsecase(c.Request().Context(), req)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusOK, "<login bSuccess='0' sMsg='The username and password you entered did not match. Please check the spelling and try again.'/>")
	}

	resp := auth.LoginResponseDTO{
		BSuccess:   1,
		SMsg:       "Login successful",
		IAccess:    user.Access,
		IUpg:       user.Upgrade,
		IAge:       user.Age,
		SToken:     user.Password,
		DUpgExp:    user.Upgrade,
		IUpgDays:   user.UpgDays,
		ISendEmail: user.EmailActive,
		StrEmail:   user.Email,
		BCCOnly:    0,
		Servers: []auth.ServerResponseDTO{
			{
				SName:   "AQW",
				SIP:     "127.0.0.1",
				ICount:  1,
				IMax:    50,
				BOnline: 1,
				BChat:   1,
				IChat:   1,
				BUpg:    0,
			},
			{
				SName:   "WQW",
				SIP:     "128.98.123",
				ICount:  1,
				IMax:    50,
				BOnline: 1,
				BChat:   1,
				IChat:   1,
				BUpg:    0,
			},
		},
	}
	return c.XML(http.StatusOK, resp)
}

func (h *Handler) RegisterUserHandler(c echo.Context) error {
	var req auth.RegisterUserRequestDTO
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return c.String(http.StatusOK, auth.RegisterUserResponseDTO{
			Status:          "Failure",
			StrErr:          "Error Code 523.02",
			StrReason:       "Bad or missing information!",
			StrButtonName:   "Back",
			StrButtonAction: "Username",
			StrMsg:          "The information you entered was rejected by the server. Please go back and make sure that you filled out everything properly.",
		}.ToString())
	}

	if req.Username == "" || req.Password == "" || (req.Gender != "M" && req.Gender != "F") || req.Age == 0 || req.DOB == "" || req.ClassID == 0 || req.EyeColor == 0 || req.HairColor == 0 || req.SkinColor == 0 || req.HairID == 0 || req.Email == "" {
		return c.String(http.StatusOK, auth.RegisterUserResponseDTO{
			Status:          "Failure",
			StrErr:          "Error Code 523.02",
			StrReason:       "Bad or missing information!",
			StrButtonName:   "Back",
			StrButtonAction: "Username",
			StrMsg:          "The information you entered was rejected by the server. Please go back and make sure that you filled out everything properly.",
		}.ToString())
	}

	userId, err := h.usecase.registerUserUsecase(c.Request().Context(), &req)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusOK, auth.RegisterUserResponseDTO{
			Status:          "Failure",
			StrErr:          "Error Code 523.14",
			StrReason:       err.Error(),
			StrButtonName:   "Back",
			StrButtonAction: "Username",
			StrMsg:          "The information you entered was rejected by the server. Please go back and make sure that you filled out everything properly.",
		}.ToString())
	}
	return c.String(http.StatusOK, auth.RegisterUserResponseDTO{
		Status: "Success",
		UserId: &userId,
		StrMsg: "Account created successfully.",
	}.ToString())
}
