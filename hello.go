package main

import "fmt"

const helloPrefixEnglish = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "

func main() {
	fmt.Println(Hello("Burak", ""))

}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingsPrefix(language) + name
}

func greetingsPrefix(language string) (prefix string) {
	prefix = helloPrefixEnglish

	switch language {
	case "spanish":
		prefix = helloPrefixSpanish

	case "french":
		prefix = helloPrefixFrench
	}

	return
}
