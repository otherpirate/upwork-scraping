package selenium_service

import (
	"fmt"
	"os"
	"time"

	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const stepWait = 2 * time.Second

type seleniumService struct {
	service   *selenium.Service
	webDriver selenium.WebDriver
}

func (s *seleniumService) Close() {
	s.webDriver.Quit()
	s.service.Stop()
}

func (s *seleniumService) Navigate(url string) error {
	err := s.webDriver.Get(url)
	if err != nil {
		return err
	}
	time.Sleep(stepWait)
	return nil
}

func (s *seleniumService) WaitElement(by, value string) (selenium.WebElement, error) {
	var err error
	var elem selenium.WebElement
	for retry := 0; retry < 5; retry++ {
		time.Sleep(stepWait)
		elem, err = s.webDriver.FindElement(by, value)
		if elem != nil {
			return elem, err
		}
	}
	return nil, err
}

func (s *seleniumService) PageSource() (string, error) {
	return s.webDriver.PageSource()
}

func NewService() (*seleniumService, error) {
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
	return &seleniumService{
		service:   service,
		webDriver: webDriver,
	}, nil
}
