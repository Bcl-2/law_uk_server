package articles

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func ReigsterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{DB: db}

	routes := r.Group("/api/v1")
    routes.GET("/articles", h.GetArticles)
	routes.GET("/chapters", h.GetChapters)
	routes.GET("/parts", h.GetParts)
	routes.GET("/sections", h.GetSections)
	routes.POST("/punishments", h.PostPunishment)
}