package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/enesyalinkaya/go-kv-store/pkg/memoryDB"
	"github.com/enesyalinkaya/go-kv-store/pkg/settings"
	"github.com/enesyalinkaya/go-kv-store/routers"
)

func main() {
	var err error

	// load settings
	settings.Setup()

	// create memoryDB client
	db := memoryDB.NewMemoryClient(settings.MemoryDBSettings.DirName, settings.MemoryDBSettings.FileName)

	// load data from file
	err = db.LoadFile()
	if err != nil {
		log.Fatal(err)
	}

	// to checking saving file
	err = db.SaveFile()
	if err != nil {
		log.Fatal(err)
	}

	// activate AutoSave operation
	db.AutoSave(time.Duration(settings.MemoryDBSettings.AutoSaveInterval))

	// build controller
	router := routers.BuildController(db)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	server := &http.Server{
		Addr:         settings.ServerSettings.Addr,
		Handler:      (logging(logger)(router)),
		ErrorLog:     logger,
		ReadTimeout:  settings.ServerSettings.ReadTimeout,
		WriteTimeout: settings.ServerSettings.WriteTimeout,
		IdleTimeout:  settings.ServerSettings.IdleTimeout,
	}

	log.Println("starting to listen on port " + settings.ServerSettings.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// logger for http requests
func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				logger.Println(r.Method, r.URL.Path, time.Since(start), r.RemoteAddr)
			}()
			next.ServeHTTP(w, r)
		})
	}
}
