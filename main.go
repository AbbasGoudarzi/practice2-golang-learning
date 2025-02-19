package main

import (
	"fmt"
	"sort"
)

var books [][]string

func main() {
	showMenu()
}

func showMenu() {
	var choiceOfUser string
	fmt.Println("Library Management Program Menu:")
	fmt.Println("1. Add a book")
	fmt.Println("2. Display books")
	fmt.Println("3. Sort books by year of publication")
	fmt.Println("4. Exit")
	fmt.Print("Please enter your choice: ")
	fmt.Scanln(&choiceOfUser)

	switch choiceOfUser {
	case "1":
		addBook()
	case "2":
		displayBook()
	case "3":
		sortBooks()
	case "4":
		return
	}
}

func addBook() {
	var title, author, yearOfPublication string
	fmt.Print("Enter the book title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter the book author: ")
	fmt.Scanln(&author)
	fmt.Print("Enter the year of publication: ")
	fmt.Scanln(&yearOfPublication)

	books = append(books, []string{title, author, yearOfPublication})

	fmt.Println("")
	showMenu()
}

func displayBook() {
	for _, book := range books {
		fmt.Println("Title:", book[0], ", Author:", book[1], ", Year of Publication:", book[2])
	}

	fmt.Println("")
	showMenu()
}

func sortBooks() {
	sort.SliceStable(books, func(i, j int) bool {
		return books[i][2] < books[j][2]
	})

	fmt.Println("Books have been sorted by year of publication")
	fmt.Println("")
	showMenu()
}
