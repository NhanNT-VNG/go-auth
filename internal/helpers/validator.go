package helpers

import (
	"fmt"
	"unicode"

	"github.com/go-playground/validator"
)

func MapValidateError(tag, fieldName, param string, input interface{}) string {
	if tag == "required" {
		return fmt.Sprintf("%s cannot be empty!", fieldName)
	}

	if tag == "min" {
		return fmt.Sprintf("%s must be at least %s characters long", fieldName, param)
	}

	if tag == "max" {
		return fmt.Sprintf("%s must be at least %s characters long", fieldName, param)
	}

	if tag == "email" {
		return fmt.Sprintf("%s is not a valid email address", input)
	}

	if tag == "strong-password" {
		return "Passwords must have at least 8 characters with at least 1 upper Case, 1 lower case, 1 numeric and 1 special character"
	}

	if tag == "eqfield" {
		return fmt.Sprintf("%s and %v does not match", param, fieldName)
	}

	return "Error validator"
}

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
