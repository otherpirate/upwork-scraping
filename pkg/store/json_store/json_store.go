package json_store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/otherpirate/upwork-scraping/pkg/models"
	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/otherpirate/upwork-scraping/pkg/utils"
)

const fileModePerm = 0755

func NewStore() *storeJSON {
	return &storeJSON{
		Path: settings.StorePath,
	}
}

type storeJSON struct {
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

func save(filePath string, obj interface{}) error {
	err := createDirectory(filePath)
	if err != nil {
		return err
	}
	json, err := utils.ToJSON(obj)
	return ioutil.WriteFile(filePath, json, fileModePerm)
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	return false
}

func loadProfile(filePath string) models.Profile {
	jsonFile, _ := os.Open(filePath)
	defer jsonFile.Close()
	jsonByte, _ := ioutil.ReadAll(jsonFile)
	profile := models.Profile{}
	json.Unmarshal(jsonByte, &profile)
	return profile
}

func (s *storeJSON) SaveProfile(profile *models.Profile) error {
	file := fmt.Sprintf("%s/profile/%s.json", s.Path, profile.ID)
	profile.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05.999999Z")
	profile.UpdatedAt = profile.CreatedAt
	status := "created"
	if fileExists(file) {
		status = "updated"
		loadedProfile := loadProfile(file)
		profile.CreatedAt = loadedProfile.CreatedAt
	}
	err := save(file, *profile)
	if err == nil {
		log.Printf("Profile %s to %s", status, file)
	} else {
		log.Printf("Could not save profile %s", profile.ID)
	}
	return err
}

func (s *storeJSON) SaveJob(name string, job *models.Job) error {
	file := fmt.Sprintf("%s/jobs/%s.json", s.Path, name)
	err := save(file, *job)
	if err == nil {
		log.Printf("Job saved to %s", file)
	} else {
		log.Printf("Could not save job %s", name)
	}
	return err
}
