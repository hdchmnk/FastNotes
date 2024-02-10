package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{Service: s}
}

// @Summary CreateUser
// @Schemes
// @Description [POST] CreateUser
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} user.User
// @Router /signup [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Login
// @Schemes
// @Description [POST] Login
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} user.LoginUserRes
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var user LoginUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", u.accessToken, 3600, "/", "localhost", false, true)

	res := &LoginUserRes{
		ID:       u.ID,
		Username: u.Username,
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Logout
// @Schemes
// @Description [GET] Logout
// @Tags user
// @Produce json
// @Success 200 {string} logout successfully
// @Router /logout [get]
func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
}
