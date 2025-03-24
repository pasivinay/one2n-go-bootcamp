package helloworld

import (
	_ "fmt"
)

const (
	spanish      = "Spanish"
	french       = "French"
	hindi        = "Hindi"
	englishHello = "Hello"
	spanishHello = "Hola"
	frenchHello  = "Bonjour"
	hindiHello   = "Namaste"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return getprefix(language) + ", " + name
}

func getprefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	case hindi:
		prefix = hindiHello
	default:
		prefix = englishHello
	}

	return
}
