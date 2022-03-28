package stub

import "fmt"

// HelloWorld prints out a lovely message!
func HelloWorld() string {
	return "Hello, world!"
}

func main() {
	fmt.Println(HelloWorld())
}
