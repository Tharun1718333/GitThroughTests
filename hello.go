package main

const EnglishGreeting = "Hello, "
const SpanishGreeting = "Hola, "

func Hello(k string, language string) string {
	if language == "" {
		if k == "" {
			return EnglishGreeting + "World"
		}
		return EnglishGreeting + k
	}
	if language == "Spanish" {
		if k == "" {
			return SpanishGreeting + "World"
		}
		return SpanishGreeting + k
	}
	return k
}
func main() {
	//fmt.Println(Hello())
}
