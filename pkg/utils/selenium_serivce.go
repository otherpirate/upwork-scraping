package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type SeleniumService struct {
	service   *selenium.Service
	WebDriver selenium.WebDriver
}

func (s *SeleniumService) Close() {
	s.WebDriver.Quit()
	s.service.Stop()
}

func (s *SeleniumService) Navigate(url string) error {
	return s.WebDriver.Get(url)
}

func (s *SeleniumService) WaitElement(by, value string) (selenium.WebElement, error) {
	var err error
	var elem selenium.WebElement
	for retry := 0; retry < 15; retry++ {
		elem, err = s.WebDriver.FindElement(by, value)
		if elem != nil {
			return elem, err
		}
		time.Sleep(1 * time.Second)
	}
	return nil, err
}

func NewSeleniumService() (*SeleniumService, error) {
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.ChromeDriver(settings.ChromeDriver),
	}
	if settings.SeleniumDebug {
		selenium.SetDebug(true)
		opts = append(opts, selenium.Output(os.Stderr))
	}
	service, err := selenium.NewSeleniumService(settings.SeleniumPath, settings.PortSelenium, opts...)
	if err != nil {
		return nil, err
	}
	configs := make(map[string]interface{})
	configs["profile.default_content_settings.popup"] = 0
	chromeCaps := chrome.Capabilities{
		Prefs: configs,
		Args:  []string{"--disable-dev-shm-usage", "--no-sandbox", "--ignore-certificate-errors", "--allow-insecure-localhost"},
	}
	caps := selenium.Capabilities{
		"browserName":         "chrome",
		"acceptInsecureCerts": true,
	}
	caps.AddChrome(chromeCaps)
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", settings.PortSelenium))
	if err != nil {
		return nil, err
	}
	return &SeleniumService{
		service:   service,
		WebDriver: webDriver,
	}, nil
}
