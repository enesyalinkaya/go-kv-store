// Package memoryDB provides in-memory key value database functionality
package memoryDB

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type MemoryClient struct {
	filePath string
	items    map[string]string
	mu       sync.RWMutex
}

// it creates memorydb client
func NewMemoryClient(dirName, fileName string) *MemoryClient {
	// create directory if it doesn't exist'
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	filePath := filepath.Join(dirName, fileName)
	return &MemoryClient{filePath: filePath, items: make(map[string]string)}
}

// Set value
func (mr *MemoryClient) Set(key string, value string) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	mr.items[key] = value
}

// Get value
func (mr *MemoryClient) Get(key string) string {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	item := mr.items[key]
	return item
}

// Flush memory
func (mr *MemoryClient) Flush() {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	mr.items = map[string]string{}
}

// load data from file
func (mr *MemoryClient) LoadFile() error {
	// check data file exists
	if _, err := os.Stat(mr.filePath); os.IsNotExist(err) {
		return nil
	}
	fp, err := os.Open(mr.filePath)
	if err != nil {
		return err
	}
	defer fp.Close()
	err = mr.load(fp)
	if err != nil {
		return err
	}
	fmt.Println("Loaded data from :", mr.filePath)
	return nil
}

func (mr *MemoryClient) load(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	items := map[string]string{}
	err = json.Unmarshal(b, &items)
	if err != nil {
		return err
	}

	mr.items = items
	return err
}

// save data to file
func (mr *MemoryClient) SaveFile() error {
	fp, err := os.Create(mr.filePath)
	if err != nil {
		return err
	}
	defer fp.Close()
	err = mr.save(fp)
	if err != nil {
		return err
	}
	fmt.Println("Saved data to :", mr.filePath, "Time:", time.Now())
	return nil
}

func (mr *MemoryClient) save(w io.Writer) (err error) {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	jsonString, err := json.Marshal(mr.items)
	_, err = w.Write(jsonString)
	return err
}

// AutoSave go routine
func (mr *MemoryClient) AutoSave(minute time.Duration) {
	ticker := time.NewTicker(minute * time.Minute)
	go func() {
		for {
			t := <-ticker.C
			err := mr.SaveFile()
			if err != nil {
				log.Fatal(err, t)
			}
		}
	}()
}
