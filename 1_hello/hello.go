package hello

const (
	french        = "French"
	spanish       = "Spanish"
	englishPrefix = "Hello "
	spanishPrefix = "Hola "
	frenchPrefix  = "Bonjour "
)

func Hello(name, language string) string {
	prefix := englishPrefix

	if name == "" {
		name = "World"
	}
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}
	return prefix + name
}
