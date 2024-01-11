package main

import (
	"fmt"
	"pustaka-api/calculation"
)

func Fundamental() {
	fmt.Println("Halo, Belajar Golang")

	result := calculation.Add(8, 9)

	fmt.Println(result)
}

func Looping() {
	for i := 0; i <= 100; i++ {
		fmt.Println("saya sendang tidur: ", i)
	}

	title := "golang is the best language"

	for index, letter := range title {
		fmt.Println("index: ", index, "letter: ", string(letter))
	}

	// for _, letter := range title {
	// 	fmt.Println("letter: ", string(letter))
	// }
}

func array() {
	// var languages [5]string
	// languages[0] = "go"
	// languages[1] = "Ruby"
	// languages[2] = "C#"
	// languages[3] = "javascript"
	// languages[4] = "python"

	languages := [...]string{
		"ruby",
		"javascript",
		"kotlin",
		"go",
		"C",
		"python",
		"java",
	}

	// languages := [5]string{"go", "ruby", "c#", "javascript", "python"}

	fmt.Println(languages)
	length := len(languages)
	fmt.Println(length)
}

func arrayLoop() {
	languages := [...]string{
		"ruby",
		"javascript",
		"kotlin",
		"go",
		"C",
		"python",
		"java",
	}

	// for index, lang := range languages {
	// 	fmt.Println("index: ", index, "language : ", lang)
	// }
	fmt.Println(languages)
}

func sliced() {
	// gamingConsole := []string{"ps4","nintendo Switch", "Xbox one"} //bentuk deklarasi lain slice

	var gamingConsole []string

	gamingConsole = append(gamingConsole, "Playstation5") //bentuk deklarasi dinamis
	gamingConsole = append(gamingConsole, "Nintendo Switch")
	gamingConsole = append(gamingConsole, "Xbox One")

	for _, console := range gamingConsole {
		fmt.Println(console)
	}
	// fmt.Println(gamingConsole)
}

func maps() {
	var MyMap map[string]int = map[string]int{}
	// MyMap = map[string]int{}

	MyMap["ruby"] = 9
	MyMap["javascript"] = 8
	MyMap["Go"] = 10

	fmt.Println(MyMap)
	fmt.Println(MyMap["ruby"]) //untuk mendapatkan nilai ruby
}

func main() {
	maps()
}