package utils

import (
	"net/mail"
	"regexp"
)

func IsValidPhone(phone string) bool {
	re := regexp.MustCompile(`^[0-9]{10}$`)
	return re.MatchString(phone)
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
