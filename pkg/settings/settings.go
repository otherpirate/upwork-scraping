package settings

import (
	"os"
	"strconv"
)

var UserName string
var Password string
var SecretAwnser string
var SeleniumPath string
var ChromeDriver string
var PortSelenium int
var SeleniumDebug bool

func LoadConfigs() {
	UserName = GetEnvDefault("USERNAME", "bobsuperworker")
	Password = GetEnvDefault("PASSWORD", "Argyleawesome123!")
	SecretAwnser = GetEnvDefault("SECRET_AWNSER", "Bobworker")
	SeleniumPath = GetEnvDefault("SELENIUM_SERVER_JAR_1", "/home/mauromurari/source/upwork-scraping/selenium_files/selenium-server.jar")
	ChromeDriver = GetEnvDefault("SELENIUM_CHROME_DRIVER_1", "/home/mauromurari/source/upwork-scraping/selenium_files/chromedriver")
	PortSelenium, _ = strconv.Atoi(GetEnvDefault("SELENIUM_PORT_1", "8099"))
	SeleniumDebug = GetEnvDefault("SELENIUM_DEBUG", "0") == "1"
}

func GetEnvDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}
