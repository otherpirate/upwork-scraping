package selenium_service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const stepWait = 5 * time.Second

type seleniumService struct {
	service   *selenium.Service
	webDriver selenium.WebDriver
}

func (s *seleniumService) Clear() {
	err := s.webDriver.DeleteAllCookies()
	if err != nil {
		log.Println("Could not clear service. Reason: ", err)
	}
	time.Sleep(stepWait)
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

func (s *seleniumService) WaitElementText(by, value, text string) (selenium.WebElement, error) {
	var err error
	var elem selenium.WebElement
	for retry := 0; retry < 5; retry++ {
		time.Sleep(stepWait)
		elems, err := s.webDriver.FindElements(by, value)
		for _, elem := range elems {
			textElement, err := elem.Text()
			if err != nil {
				continue
			}
			if textElement == text {
				return elem, nil
			}
		}
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
	configs["browser.cache.disk.enable"] = 0
	configs["browser.cache.memory.enable"] = 1
	configs["browser.cache.offline.enable"] = 0
	configs["network.http.use-cache"] = 1
	chromeCaps := chrome.Capabilities{
		Prefs: configs,
		Args: []string{
			"--disable-dev-shm-usage",
			"--ignore-certificate-errors",
			"--allow-insecure-localhost",
			"--disable-back-forward-cache",
			"--enable-javascript",
		},
	}
	caps := selenium.Capabilities{
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
