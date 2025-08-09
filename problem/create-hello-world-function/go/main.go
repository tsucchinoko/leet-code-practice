package main

func createHelloWorld(args ...any) func() string {
	return func() string {
		return "Hello World!"
	}
}
