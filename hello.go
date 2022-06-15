package main

const EnglishGreeeting = "Hello, "

func Hello(k string) string {
	if k == "" {
		return EnglishGreeeting + "World"
	}
	return EnglishGreeeting + k
}
func main() {
	//fmt.Println(Hello())
}
