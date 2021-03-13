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
var RabbitQueueRetry string
var RabbitQueueFailure string
var CrawlerMaxRetries int64
var WaitBeforeRequeue int

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
	RabbitQueueRetry = GetEnvDefault("RABBIT_QUEUE_USER", "upwork-scraping-retry")
	RabbitQueueFailure = GetEnvDefault("RABBIT_QUEUE_USER", "upwork-scraping-failure")
	CrawlerMaxRetries, _ = strconv.ParseInt(GetEnvDefault("CRAWLER_MAX_RETRIES", "10"), 10, 64)
	WaitBeforeRequeue, _ = strconv.Atoi(GetEnvDefault("WAIT_BEFORE_REQUEUE", "5"))
}

func GetEnvDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}
