package helloworld

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}
	return langPrefix(lang) + name
}

func langPrefix(lang string) string {
	prefix := englishHelloPrefix
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix
}
