package main

func CreateHelloWorld(args ...any) func() string {
	return func() string {
		return "Hello World!"
	}
}
