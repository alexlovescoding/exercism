package hello

const TestVersion = 1

func HelloWorld(name string) string {
	if len(name) == 0 {
		return "Hello, World!"
	}
	return "Hello, " + name + "!"
}
