// settings package provides get settings from environment variables.
package settings

import (
	"net"
	"os"
	"strconv"
	"time"
)

type Server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

var ServerSettings = &Server{}

type MemoryDB struct {
	DirName          string
	FileName         string
	AutoSaveInterval int
}

var MemoryDBSettings = &MemoryDB{}

func Setup() {
	serverSettingsLoad()
	memorySettingsLoad()
}

// load settings from environment variable
func serverSettingsLoad() {
	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	addr := ""
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}

	readTimeout := 60
	if v, err := strconv.Atoi(os.Getenv("READ_TIMEOUT")); err != nil && v != 0 {
		readTimeout = v
	}

	writeTimeout := 60
	if v, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT")); err != nil && v != 0 {
		writeTimeout = v
	}

	idleTimeout := 10
	if v, err := strconv.Atoi(os.Getenv("IDLE_TIMEOUT")); err != nil && v != 0 {
		idleTimeout = v
	}

	ServerSettings.Addr = net.JoinHostPort(addr, port)
	ServerSettings.ReadTimeout = time.Duration(readTimeout) * 1000000000
	ServerSettings.WriteTimeout = time.Duration(writeTimeout) * 1000000000
	ServerSettings.IdleTimeout = time.Duration(idleTimeout) * 1000000000
}

func memorySettingsLoad() {
	dirName := "tmp"
	if v := os.Getenv("DIR_NAME"); v != "" {
		dirName = v
	}

	fileName := "data.txt"
	if v := os.Getenv("FILE_NAME"); v != "" {
		fileName = v + ".txt"
	}

	autoSaveInterval := 1
	if v, err := strconv.Atoi(os.Getenv("AUTO_SAVE_INTERVAL")); err != nil && v != 0 {
		autoSaveInterval = v
	}
	MemoryDBSettings.DirName = dirName
	MemoryDBSettings.FileName = fileName
	MemoryDBSettings.AutoSaveInterval = autoSaveInterval

}
