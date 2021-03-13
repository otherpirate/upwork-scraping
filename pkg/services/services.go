package services

import (
	"github.com/tebeka/selenium"
)

type Service interface {
	Close()
	Navigate(url string) error
	WaitElement(by, value string) (selenium.WebElement, error)
	WaitElementText(by, value, text string) (selenium.WebElement, error)
	PageSource() (string, error)
	Clear()
}
