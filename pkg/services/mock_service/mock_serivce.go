package mock_service

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/tebeka/selenium"
)

type mockService struct {
	url          string
	baseHTMLPath string
}

func (s *mockService) Clear() {
}

func (s *mockService) Close() {
}

func (s *mockService) Navigate(url string) error {
	s.url = url
	return nil
}

func (s *mockService) WaitElement(by, value string) (selenium.WebElement, error) {
	return mockedElement{}, nil
}

func (s *mockService) WaitElementText(by, value, text string) (selenium.WebElement, error) {
	return mockedElement{}, nil
}

func (s *mockService) PageSource() (string, error) {
	path := strings.ReplaceAll(s.url, "https://www.upwork.com/", "")
	path = strings.ReplaceAll(path, "/", "_")
	path = fmt.Sprintf("%s/html_pages/%s.html", s.baseHTMLPath, path)
	path, _ = filepath.Abs(path)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func NewService() (*mockService, error) {
	return NewServicePath("..")
}

func NewServicePath(baseHTMLPath string) (*mockService, error) {
	return &mockService{
		baseHTMLPath: baseHTMLPath,
	}, nil
}

type mockedElement struct {
	selenium.WebElement
}

func (m mockedElement) Click() error {
	return nil
}

func (m mockedElement) SendKeys(keys string) error {
	return nil
}
