package assets

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/negrel/assert"
)

//go:embed public/*
var assets embed.FS

func disableCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func AssetsHandler(isDevelopment bool) http.Handler {
	if isDevelopment == true {
		workDir, _ := os.Getwd()
		root := http.Dir(filepath.Join(workDir, "internal/server/assets/public"))
		return disableCache(http.FileServer(root))
	}

	publicFs, err := fs.Sub(assets, "public")
	assert.NoError(err)

	fs := http.FileServer(http.FS(publicFs))
	return fs
}
