package TestLibrary

import (
	"strings"
	"unicode"
)

// ToUpper convierte una cadena a mayúsculas
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower convierte una cadena a minúsculas
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Capitalize capitaliza la primera letra de cada palabra
func Capitalize(s string) string {
	return strings.Title(s)
}

// CapitalizeWords capitaliza la primera letra de cada palabra usando Runas
func CapitalizeWords(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return r
		}
		return unicode.ToUpper(r)
	}, s)
}
