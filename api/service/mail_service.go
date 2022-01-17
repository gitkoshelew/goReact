package service

import (
	"errors"
	"fmt"
	"goReact/webapp"
	"goReact/webapp/server/logging"
	"net/smtp"
	"os"
)

// MailType ...
type MailType string

// List of Mail Types
const (
	EmailConfirmation MailType = "confirmation"
	PassReset         MailType = "passReset"
)

// Mail ...
type Mail struct {
	Source   string
	Address  string
	From     string
	Password string
	Logger   *logging.Logger
}

// GetMail ...
func GetMail(config *webapp.Config) *Mail {
	return &Mail{
		Source:   config.MailService.Source,
		Address:  config.MailService.Address,
		From:     config.MailService.From,
		Password: config.MailService.Password,
		Logger:   logging.GetLogger(),
	}
}

// Auth ...
func (m *Mail) Auth() smtp.Auth {
	return smtp.PlainAuth("", m.From, m.Password, m.Source)
}

// Create ...
func (m *Mail) Create(mailType MailType, message string, to []string) (string, error) {

	switch mailType {
	case EmailConfirmation:
		return "From: " + m.From + "\n" +
				"To: " + to[0] + "\n" +
				"Subject: Email confirmation\n\n" +
				"Dear client! Thank you for choising our Hotel, we will look after your pets like our own!\n" +
				"To complete your registration and activate your account, simply verify your email address by link below: \n" +
				fmt.Sprintf("%s%s \n", os.Getenv("EMAIL_CONFITM_ENDPOINT"), message) +
				"Link will be expire in 2 hours.",
			nil
	case PassReset:
		return "From: " + m.From + "\n" +
				"To: " + to[0] + "\n" +
				"Subject: Password reset\n\n" +
				fmt.Sprintf("%s%s \n", os.Getenv("RESTORE_PASSWORD_ENDPOINT"), message) +
				"Link will be expire in 1 hours.",
			nil
	default:
		return "", errors.New("Ivalid mail type")
	}
}

// Send ...
func (m *Mail) Send(mailType MailType, messege string, to []string) error {

	msg, err := m.Create(mailType, messege, to)
	if err != nil {
		m.Logger.Errorf("Mail message creation error: %v", err)
		return err
	}

	err = smtp.SendMail(
		m.Address,
		m.Auth(),
		m.From,
		to,
		[]byte(msg))
	if err != nil {
		m.Logger.Errorf("Mail sending error: %v", err)
		return err
	}

	m.Logger.Info("Mail send successfull")
	return nil
}
