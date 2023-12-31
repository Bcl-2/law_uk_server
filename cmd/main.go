package main

import (
	"main/pkg/articles"
	"main/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("../pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

    
	articles.ReigsterRoutes(r, h)

    r.Run(port)
}