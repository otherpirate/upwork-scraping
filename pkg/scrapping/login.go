package scrapping

import (
	"log"
)

func (u *Upwork) login(userName, password, secretAwnser string) error {
	log.Println("Login into portal...")
	err := u.service.Navigate("https://www.upwork.com/ab/account-security/login")
	if err != nil {
		return err
	}
	err = u.fillUser(userName)
	if err != nil {
		return err
	}
	err = u.fillPassword(password)
	if err != nil {
		return err
	}
	err = u.fillSecretAwnser(secretAwnser)
	if err != nil {
		return nil
	}
	log.Println("Login done")
	return nil
}

func (u *Upwork) fillUser(userName string) error {
	userInput, err := u.service.WaitElement("id", "login_username")
	if err != nil {
		return err
	}
	userInput.SendKeys(userName)
	continueButton, err := u.service.WaitElement("id", "login_password_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}

func (u *Upwork) fillPassword(password string) error {
	passInput, err := u.service.WaitElement("id", "login_password")
	if err != nil {
		return err
	}
	passInput.SendKeys(password)
	continueButton, err := u.service.WaitElement("id", "login_control_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}

func (u *Upwork) fillSecretAwnser(secretAwnser string) error {
	secretAwserInput, err := u.service.WaitElement("id", "login_deviceAuthorization_answer")
	if err != nil {
		return nil
	}
	secretAwserInput.SendKeys(secretAwnser)
	continueButton, err := u.service.WaitElement("id", "login_control_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}
