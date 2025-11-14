package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"time"
)

func SendEmailOTP(to, otp string) error {

	smtpHost := "smtpout.secureserver.net"
	smtpPort := "587"

	smtpUser := "kartikk@porcupi9.com"
	smtpPass := "Kartik2345"

	displayFrom := "no-reply@porcupi9.com"

	subject := "Your OTP Code"
	body := fmt.Sprintf("Your OTP code is: %s\nThis OTP will expire in 5 minutes.", otp)

	msg := []byte("To: " + to + "\r\n" +
		"From: " + displayFrom + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Reply-To: " + displayFrom + "\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n\r\n" +
		body + "\r\n")

	// Faster connection handling
	dialer := net.Dialer{Timeout: 5 * time.Second}

	log.Println("[SMTP] Connecting to:", smtpHost+":"+smtpPort)

	conn, err := dialer.Dial("tcp", smtpHost+":"+smtpPort)
	if err != nil {
		log.Println("[SMTP ERROR] TCP connection failed:", err)
		return err
	}
	log.Println("[SMTP] Connected")

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Println("[SMTP ERROR] Client creation failed:", err)
		return err
	}

	tlsConfig := &tls.Config{ServerName: smtpHost}
	if err = c.StartTLS(tlsConfig); err != nil {
		log.Println("[SMTP ERROR] STARTTLS failed:", err)
		return err
	}

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	if err = c.Auth(auth); err != nil {
		log.Println("[SMTP ERROR] Auth failed:", err)
		return err
	}

	if err = c.Mail(smtpUser); err != nil {
		log.Println("[SMTP ERROR] MAIL FROM failed:", err)
		return err
	}

	if err = c.Rcpt(to); err != nil {
		log.Println("[SMTP ERROR] RCPT failed:", err)
		return err
	}

	w, err := c.Data()
	if err != nil {
		log.Println("[SMTP ERROR] Data init failed:", err)
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		log.Println("[SMTP ERROR] Writing failed:", err)
		return err
	}

	_ = w.Close()
	log.Println("[SMTP] Email sent")

	return c.Quit()
}
