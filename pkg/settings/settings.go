package settings

import (
	"os"
	"path/filepath"
	"strconv"
)

var SeleniumPath string
var ChromeDriver string
var PortSelenium int
var SeleniumDebug bool
var StorePath string
var RabbitURI string
var RabbitQueueUser string
var RabbitQueueProfile string

func LoadConfigs() {
	basePath, _ := filepath.Abs("../")
	SeleniumPath = GetEnvDefault("SELENIUM_JAR", basePath+"/selenium_files/selenium-server.jar")
	ChromeDriver = GetEnvDefault("SELENIUM_DRIVER", basePath+"/selenium_files/chromedriver")
	PortSelenium, _ = strconv.Atoi(GetEnvDefault("SELENIUM_PORT_1", "8099"))
	SeleniumDebug = GetEnvDefault("SELENIUM_DEBUG", "0") == "1"
	StorePath = GetEnvDefault("STORE_PATH", "/tmp/upwork-scrapping/store_json")
	RabbitURI = GetEnvDefault("RABBIT_URI", "amqp://guest:guest@localhost:5672/")
	RabbitQueueUser = GetEnvDefault("RABBIT_QUEUE_USER", "upwork-scraping-user")
	RabbitQueueProfile = GetEnvDefault("RABBIT_QUEUE_USER", "upwork-scraping-profile")
}

func GetEnvDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}
