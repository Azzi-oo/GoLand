package main

import "fmt"

func main() {
	message, _ := enterTheClub(15)
	fmt.Println(message)
	// fmt.Println(entered)
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "Входи", true
	} else {
		return "Тебе нет 18", true
	}

}
