package settings

import (
	"os"
	"path/filepath"
	"strconv"
)

var UserName string
var Password string
var SecretAwnser string
var SeleniumPath string
var ChromeDriver string
var PortSelenium int
var SeleniumDebug bool
var StorePath string

func LoadConfigs() {
	basePath, _ := filepath.Abs("../")
	UserName = GetEnvDefault("USERNAME", "bobsuperworker")
	Password = GetEnvDefault("PASSWORD", "Argyleawesome123!")
	SecretAwnser = GetEnvDefault("SECRET_AWNSER", "Bobworker")
	SeleniumPath = GetEnvDefault("SELENIUM_JAR", basePath+"/selenium_files/selenium-server.jar")
	ChromeDriver = GetEnvDefault("SELENIUM_DRIVER", basePath+"/selenium_files/chromedriver")
	PortSelenium, _ = strconv.Atoi(GetEnvDefault("SELENIUM_PORT_1", "8099"))
	SeleniumDebug = GetEnvDefault("SELENIUM_DEBUG", "0") == "1"
	StorePath = GetEnvDefault("STORE_PATH", "/tmp/upwork-scrapping/store_json")
}

func GetEnvDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}
