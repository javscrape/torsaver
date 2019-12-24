package torsaver

import (
	"github.com/goextension/log/zap"
)

func init() {
	zap.DefaultZapFilePath = "tor.log"
	zap.InitZapFileSugar()
}
