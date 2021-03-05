package scrapping

import (
	"log"
)

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
		return nil
	}
	log.Println("Login done")
	return nil
}

func (u *Upwork) fillUser() error {
	userInput, err := u.service.WaitElement("id", "login_username")
	if err != nil {
		return err
	}
	userInput.SendKeys(u.userName)
	continueButton, err := u.service.WaitElement("id", "login_password_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}

func (u *Upwork) fillPassword() error {
	passInput, err := u.service.WaitElement("id", "login_password")
	if err != nil {
		return err
	}
	passInput.SendKeys(u.password)
	continueButton, err := u.service.WaitElement("id", "login_control_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}

func (u *Upwork) fillSecretAwnser() error {
	secretAwserInput, err := u.service.WaitElement("id", "login_deviceAuthorization_answer")
	if err != nil {
		return nil
	}
	secretAwserInput.SendKeys(u.secretAwnser)
	continueButton, err := u.service.WaitElement("id", "login_control_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}
