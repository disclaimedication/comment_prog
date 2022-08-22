package main

import "fmt"

func Run() error {
	fmt.Println("This is the error handleing function")
	return nil
}
func main() {
	fmt.Println("this is the main function")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
