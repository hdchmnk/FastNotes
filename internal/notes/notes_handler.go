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

// @Summary CreateNote
// @Schemes
// @Description [POST] CreateNote
// @Tags note
// @Accept json
// @Produce json
// @Success 200 {object} notes.Note
// @Router /createnote [post]
func (h *Handler) CreateNote(c *gin.Context) {
	var n CreateNoteReq
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"CreateNote error": err.Error()})
		return
	}

	res, err := h.Service.CreateNote(c.Request.Context(), &n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

// @Summary GetNotesByUserID
// @Schemes
// @Description [POST] GetNotesByUserID
// @Tags note
// @Accept json
// @Produce json
// @Success 200 {object} []Note
// @Router /getnotesbyid [post]
func (h *Handler) GetNotesByUserID(c *gin.Context) {
	var n GetNotesByUserIDReq
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.GetNotesByUserID(c.Request.Context(), n.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"GetNotesByUserID error": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}
