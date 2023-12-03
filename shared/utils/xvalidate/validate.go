package xvalidate

import "regexp"

func IsEmail(email string) bool {
	pattern := `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

func IsMobile(mobile string) bool {
	return false
}
