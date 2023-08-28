package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/config"
	"github.com/sunflower10086/TikTok/http/internal/dao/db"
	___interaction "github.com/sunflower10086/TikTok/http/internal/interaction/rpc"
	___relation "github.com/sunflower10086/TikTok/http/internal/relation/rpc"
	"github.com/sunflower10086/TikTok/http/internal/router"
)

var configFile = flag.String("f", "etc/", "the config file")

func init() {
	config.LoadConfigFromYaml(*configFile)
	db.Init()
	go ___interaction.Init()
	go ___relation.Init()
}

func main() {
	r := gin.Default()

	router.Init(r)

	httpConf := config.C().Apps.HTTP
	addr := fmt.Sprintf("%s:%s", httpConf.Host, httpConf.Port)
	r.Run(addr)

}
