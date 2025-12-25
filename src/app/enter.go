package app

import (
	"flag"

	"github.com/wxq/metaland-blog/src/config"
	"github.com/wxq/metaland-blog/src/db"
	"github.com/wxq/metaland-blog/src/router"
	"github.com/wxq/metaland-blog/src/xzap"
	"github.com/wxq/metaland-blog/src/xzap/logger"
)

const (
	DefaultConfigPath = "./conf/config.yaml"
)

func Start() {
	conf := flag.String("conf", DefaultConfigPath, "conf file path")
	flag.Parse()
	c := config.LoadConfig(*conf)

	xzap.Initialize(c)
	if err := db.Initialize(c); err != nil {
		logger.Error(err.Error())
		return
	}
	router.Start()
}
