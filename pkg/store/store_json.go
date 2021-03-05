package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/otherpirate/upwork-scraping/pkg/models"
)

const fileModePerm = 0644

type StoreJSON struct {
	Path string
}

func createDirectory(path string) error {
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, fileModePerm)
	}
	return err
}

func (s *StoreJSON) save(filePath string, obj interface{}) error {
	err := createDirectory(filePath)
	if err != nil {
		return err
	}
	json, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, json, fileModePerm)
}

func (s *StoreJSON) SaveJob(name string, obj models.Job) error {
	file := fmt.Sprintf("%s/jobs/%s.json", s.Path, name)
	err := s.save(file, obj)
	if err == nil {
		log.Printf("Job saved to %s", file)
	} else {
		log.Printf("Could not save job %s", name)
	}
	return err
}
