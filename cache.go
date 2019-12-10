package torsaver

import (
	"crypto/sha256"
	"fmt"
	cache "github.com/gocacher/badger-cache/v2"
	"github.com/gocacher/cacher"
	"io/ioutil"
	"net/http"
)

var useCache bool = false

func RegisterCache() {
	cacher.Register(cache.New())
	useCache = true
}

func Get(url string) (e error) {
	name := Hash(url)
	get, e := http.Get(url)
	if e != nil {
		return e
	}
	bys, e := ioutil.ReadAll(get.Body)

	err := cacher.Set(name, bys)
	if err != nil {
		return
	}
}

// Hash ...
func Hash(url string) string {
	sum256 := sha256.Sum256([]byte(url))
	return fmt.Sprintf("%x", sum256)
}
