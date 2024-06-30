package validator

import (
	"errors"
	"regexp"
	"time"
)

func ValidatePhone(phone string) bool {
	return regexp.MustCompile(`^\+998[0-9]{9}$`).MatchString(phone)
}

func ValidateGmail(gmail string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail.com$`).MatchString(gmail)
}

func ValidateGender(gender string) bool {
	if gender!="male" && gender!="female" && gender!="other" {
		return false
	}
	return true
}

func ValidateDaySunDay(day string)  error {
	date, err := time.Parse("2006-01-02", day)
	if err != nil {
		return err
	}

	if date.Weekday() != time.Sunday {
		return errors.New("wrong day: valid only Sunday")
	}

	return nil
}