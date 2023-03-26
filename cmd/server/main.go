package main

import "fmt"

// Run - responsible for installation and starting of app
func Run() error {
	fmt.Println("Staring app")
	return nil
}

func main() {
	fmt.Println("Rest Api ")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
