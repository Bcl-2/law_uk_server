package articles

import (
	"main/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetChapters(c *gin.Context) { 
	var chapters []models.Article_Chapters
	if result := h.DB.Find(&chapters); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	c.JSON(http.StatusOK, &chapters)
}