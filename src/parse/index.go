package parse

import (
	"github.com/sudo-at-night/me-go/src/cache"

	"errors"
	"log"
	"path/filepath"
	"strings"
)

// GetIndexContent retrieves the final index.html content to be returned to the user.
func GetIndexContent(path string) (string, error) {
	cacheManager := cache.GetManager()
	meta, err := GetMeta(path)
	if err != nil {
		if errors.Is(err, errMeta404) {
			return GetIndexContent("/")
		} else if errors.Is(err, errMetaFatal404) {
			log.Fatal("Cannot fetch any meta data")
		}
	}
	index, err := cacheManager.GetFile("index")
	if err != nil {
		log.Fatalf("Cannot get index from cache: %s", err)
	}
	newHTML := strings.Replace(string(index), "<!-- SSR-HEAD -->", meta, 1)
	return newHTML, nil
}

// CacheIndex caches the initial index.[hash].html
func CacheIndex() {
	cacheManager := cache.GetManager()
	match, err := filepath.Glob("pages/index.*.html")
	if err != nil {
		log.Fatal(err)
	}
	if len(match) < 1 {
		log.Fatal("No pages/index.[hash].html template found")
	}
	if cacheManager.CacheFile(match[0], "index"); err != nil {
		log.Fatalf("Could not properly load and cache pages/index.[hash].html: %s", err)
	}
}
