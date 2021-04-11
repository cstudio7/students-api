package main

import "fmt"

func Run() error {
	fmt.Println("Running App")
	return nil
}

func main() {
	err := Run()
	if err != nil {
		fmt.Println("Error running app")
		fmt.Println(err)
	}
}
