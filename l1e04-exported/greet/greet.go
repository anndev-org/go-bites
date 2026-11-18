package greet

func Hello() string {
	return "Hello, " + secret()
}

func secret() string { return "Go" }
