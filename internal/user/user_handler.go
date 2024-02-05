package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{Service: s}
}

// @Summary CreateUser
// @Tags auth
// @Description create user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body CreateUserReq
// @Success 200 {integer} integer 1
// @Failure 500 error
// @Router /signup [post]
func (h *Handler) CreateUser(c echo.Context) error {
	var u CreateUserReq
	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	res, err := h.Service.CreateUser(c.Request().Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}

	c.JSON(http.StatusOK, res)
	return nil
}

func (h *Handler) Login(c echo.Context) error {
	var user LoginUserReq
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	u, err := h.Service.Login(c.Request().Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	cookies := new(http.Cookie)
	cookies.Name = "jwt"
	cookies.Value = u.accessToken
	cookies.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookies)

	res := &LoginUserRes{
		ID:       u.ID,
		Username: u.Username,
	}

	c.JSON(http.StatusOK, res)
	return nil
}

func (h *Handler) Logout(c echo.Context) error {
	cookies := new(http.Cookie)
	cookies.Name = "jwt"
	cookies.Value = ""
	cookies.Expires = time.Now().Add(-1 * time.Hour)
	c.JSON(http.StatusOK, "logout successfully")
	return nil
}
