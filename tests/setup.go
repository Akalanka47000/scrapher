package tests

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"scrapher/src/config"
	"scrapher/src/global"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

// ServeDirectory serves static files from the specified directory
// and listens on the given port. The directory is relative to the root of the project.
func ServeDirectory(dir string, port int) (*http.Server, string) {
	dir = fmt.Sprintf("%s/%s", basepath, dir)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set(global.HdrContentType, "text/html; charset=utf-8")
			http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
		}),
	}

	address := fmt.Sprintf("http://localhost:%d", port)

	go func() {
		log.Printf("Serving %s on %s", dir, address)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	return server, address
}

// Initializes the configuration for the tests.
func Setup() {
	config.Load()
}
