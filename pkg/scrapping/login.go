package scrapping

import (
	"log"
	"time"
)

const stepWait = 2 * time.Second

func (u *Upwork) Login() error {
	log.Println("Login into portal...")
	err := u.service.Navigate("https://www.upwork.com/ab/account-security/login")
	if err != nil {
		return err
	}
	err = u.fillUser()
	if err != nil {
		return err
	}
	err = u.fillPassword()
	if err != nil {
		return err
	}
	err = u.fillSecretAwnser()
	if err != nil {
		return err
	}
	log.Println("Login done")
	log.Println(u.service.WebDriver.PageSource())
	return nil
}

func (u *Upwork) fillUser() error {
	time.Sleep(stepWait)
	userInput, err := u.service.WaitElement("id", "login_username")
	if err != nil {
		return err
	}
	userInput.SendKeys(u.userName)
	time.Sleep(stepWait)
	continueButton, err := u.service.WaitElement("id", "login_password_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}

func (u *Upwork) fillPassword() error {
	time.Sleep(2 * time.Second)
	passInput, err := u.service.WaitElement("id", "login_password")
	if err != nil {
		return err
	}
	passInput.SendKeys(u.password)
	time.Sleep(stepWait)
	continueButton, err := u.service.WaitElement("id", "login_control_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}

func (u *Upwork) fillSecretAwnser() error {
	time.Sleep(stepWait)
	secretAwserInput, err := u.service.WaitElement("id", "login_deviceAuthorization_answer")
	if err != nil {
		return nil
	}
	secretAwserInput.SendKeys(u.secretAwnser)
	time.Sleep(stepWait)
	continueButton, err := u.service.WaitElement("id", "login_control_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}
