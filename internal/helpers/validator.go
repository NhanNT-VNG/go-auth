package helpers

import (
	"unicode"

	"github.com/go-playground/validator"
)

func StrongPassword(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	isStrongPassword := verifyPassword(value)
	return isStrongPassword

}

func verifyPassword(s string) bool {
	var (
		hasMinLen  = len(s) >= 8
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		case unicode.IsLower(c):
			hasLower = true
		default:
			// return false, false, false, false
		}
	}

	return hasMinLen && hasLower && hasUpper && hasSpecial && hasNumber
}
