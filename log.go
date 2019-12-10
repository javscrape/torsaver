package torsaver

import (
	"github.com/goextension/log"
	"github.com/goextension/log/zap"
)

func init() {
	zap.DefaultZapFilePath = "tor.log"

	log.Register(zap.NewZapFileSugar())
}
