package messages

import "fmt"

func SomethingGoesWrong() []byte {
	return []byte("Something goes wrong! :(")
}
func MethodsAllowed(methods ...string) []byte {
	return []byte(fmt.Sprint("Methods allowed:", methods))
}

func ContentTypesSupported(contentTypes ...string) []byte {
	return []byte(fmt.Sprint("Content types supported:", contentTypes))
}
