package articles

import (
	"fmt"
	"main/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Section_number int `json:"section_number"`
	Chapter_number int `json:"chapter_number"`
	Article_number float64 `json:"article_number"`
	Part_uuid string `json:"part_uuid"`
}



func (h *handler) PostPunishment(c *gin.Context) { 
	var punishments []models.Article_punishments
	var request Request
	if err := c.ShouldBindJSON(&request); err!= nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

	fmt.Println(request)


	if result := h.DB.Where(models.Article_punishments{
		Section_number: request.Section_number,
		Chapter_number: request.Chapter_number,
        Article_number: request.Article_number,
        Part_uid: request.Part_uuid}).Preload("Section_ref").
									  Preload("Chapter_ref").
									  Preload("Article_ref").
									  Preload("Part_ref").
									  Preload("Punishment_ref").
									  Preload("Punishment_min_unit_ref").
									  Preload("Punishment_max_unit_ref").
									  Preload("Danger_ref").Find(&punishments); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	c.JSON(http.StatusOK, &punishments)
}