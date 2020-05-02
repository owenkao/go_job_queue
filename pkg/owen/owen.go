package owen

import "fmt"

// Init is a test api
func Init() {
	fmt.Println("test")
}

func init() {
	fmt.Println("world1 module init function")
}
