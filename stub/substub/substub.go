package substub

import "fmt"

// HelloName prints out a greeting to the specified person!
func HelloName(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Println(HelloName("wilhelm"))
}
