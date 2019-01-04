package util

import (
	"regexp"
)

const (
	regPhone    = "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	regEmail    = "^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,4})$"
	regUserName = "^([A-Za-z_])+\\w"
)

var (
	phoneRegex    = regexp.MustCompile(regPhone)
	usernameRegex = regexp.MustCompile(regUserName)
)

func RegPhone(phone string) bool {
	return phoneRegex.MatchString(phone)
}

func RegUserName(username string) bool {
	return usernameRegex.MatchString(username)
}
