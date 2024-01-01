package articles

import (
	"main/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetParts(c *gin.Context) { 
	var parts []models.Article_parts
	if result := h.DB.Find(&parts); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	c.JSON(http.StatusOK, &parts)
}