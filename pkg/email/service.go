package email

import (
	"fmt"
	"net/smtp"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Service struct {
	config Config
}

func NewEmailService(config Config) *Service {
	return &Service{config: config}
}

func (s *Service) SendOrderConfirmation(to string, orderID uint) error {
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)

	subject := "Pedido Confirmado"
	body := fmt.Sprintf("Seu pedido #%d foi confirmado e está pago!", orderID)
	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", s.config.From, to, subject, body)

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	return smtp.SendMail(addr, auth, s.config.From, []string{to}, []byte(msg))
}
