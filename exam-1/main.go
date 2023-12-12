package main

import "fmt"

func main() {
	// Got Username and Password from keyboard
	var username string
	var password string
	
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	
	// Call function LogIn
	LogIn(username, password)
	
	// Call function Add and print the result
	result := Add(4, 5)
	fmt.Println("ผลบวกเท่ากับ:", result)

	// Call function Minus and print the result
	fmt.Print("ผลลบเท่ากับ:", Minus(3, 2))
}

func HelloFunction() string {
	return "Hello"
}

func Add(a int, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}

func UserLogin(username string, password string) bool {
	if username == "admin" && password == "password" {
		return true
	} else {
		return false
	}
}

// Use function UserLogin
func LogIn(username string, password string) {
	if UserLogin(username, password) {
		fmt.Println("เข้าสู่ระบบสำเร็จ")
	} else {
		fmt.Println("เข้าสู่ระบบผิดพลาด")
	}
}
