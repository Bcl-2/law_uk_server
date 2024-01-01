package articles

import (
	"main/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetSections(c *gin.Context) { 
	var sections []models.Article_sections
	if result := h.DB.Find(&sections); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	c.JSON(http.StatusOK, &sections)
}