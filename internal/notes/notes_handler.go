package notes

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

func (h *Handler) CreateNote(c *gin.Context) {
	var n CreateNoteReq
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateNote(c.Request.Context(), &n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetNotesByUserID(c *gin.Context) {
	var n int64
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.GetNotesByUserID(c.Request.Context(), n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}
