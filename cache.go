package torsaver

import (
	cache "github.com/gocacher/badger-cache/v2"
	"github.com/gocacher/cacher"
)

func RegisterCache() {
	cacher.Register(cache.New())
}
