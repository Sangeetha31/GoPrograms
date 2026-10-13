package main

import "fmt"

func getAppInfo() {
	fmt.Println("WebAPP2.0")
}

func printMessage(msg string) {
	fmt.Println("Message Received:", msg)
}

func getDeveloperInfo() string {
	return "Developed By Steve"
}

func main() {
	getAppInfo()
	printMessage("Welcome to my WebAPP2.0")
	dev := getDeveloperInfo()
	fmt.Println("Developer Details Here:", dev)

}
