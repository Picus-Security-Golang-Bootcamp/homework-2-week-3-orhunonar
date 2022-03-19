package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:] //Read the commands on the terminal line

	listofbooks := []books{{ // List of the Books
		Bookid:    "1",
		Bookname:  "Nutuk",
		Page:      "154",
		Stock:     10,
		Price:     "20",
		StockCode: "5555",
		ISBN:      "933-3-3",
		Author: Author{
			ID:   "123444",
			Name: "Mustafa Kemal Atatürk",
		},
	},
		{Bookid: "2",
			Bookname:  "Olasılıksız",
			Page:      "200",
			Stock:     10,
			Price:     "20",
			StockCode: "4235",
			ISBN:      "345-5-6",
			Author: Author{
				ID:   "2534563",
				Name: "Adam Fawer",
			}},
		{Bookid: "3",
			Bookname:  "Silmarillion",
			Page:      "654",
			Stock:     10,
			Price:     "20",
			StockCode: "3454",
			ISBN:      "234-5-6",
			Author: Author{
				ID:   "34563",
				Name: "J. R. R. Tolkien",
			}},
	}

	if args[0] == "list" { //Check if command is list

		for i := 0; i < len(listofbooks); i++ {
			fmt.Println(listofbooks[i].Bookname)
		}
	} else if args[0] == "search" { //Check if command is search
		for i := 0; i < len(listofbooks); i++ {
			if args[1] == listofbooks[i].Bookname { //Print the name of the books
				fmt.Println(listofbooks[i])
				break

			} else if i == len(listofbooks)-1 {

				fmt.Println("Invalid Book Name")
				break

			}
		}

	} else if args[0] == "get" { // //Check if command is get
		for i := 0; i < len(listofbooks); i++ {
			if args[1] == listofbooks[i].Bookid {
				fmt.Println(listofbooks[i])
				break

			} else if i == len(listofbooks)-1 {

				fmt.Println("Invalid Book ID")
				break

			}
		}
	} else if args[0] == "delete" { //Check if command is delete
		for i := 0; i < len(listofbooks); i++ {
			if args[1] == listofbooks[i].Bookid {
				listofbooks = remove(listofbooks, i)
				fmt.Println(listofbooks)
				break
			} else if i == len(listofbooks)-1 {

				fmt.Println("Invalid Book ID")
				break

			}
		}

	} else if args[0] == "buy" { //Check if command is buy
		for i := 0; i < len(listofbooks); i++ {
			if args[1] == listofbooks[i].Bookid {
				quantitiy, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Println(err)

				} else if listofbooks[i].Stock > quantitiy {

					listofbooks[i].Stock -= quantitiy
					fmt.Println("New quantity of ", listofbooks[i].Bookname, " is ", listofbooks[i].Stock)
					fmt.Println(listofbooks[i])
					break
				} else {
					fmt.Println("Insufficent Number of Books")
					break
				}

			} else if i == len(listofbooks)-1 {

				fmt.Println("Invalid Book ID")
				break

			}

		}

	} else { //If commands are not valid. Show the usage of the program.
		fmt.Println("\n" + "Invalid command\n" + "Usage:\n" + "list: list the books in the list\n" +
			"search <bookName>: search the name of the book if it is in the list\n" + "get <bookID>: get the details of the book that have this ID\n" +
			"delete <bookID>: delete the book that have this ID\n" + "buy <bookID> <quantity>: buy from the book with the specified id as much as the number written ")
	}

}

type books struct {
	Bookid    string
	Bookname  string
	Page      string
	Stock     int
	Price     string
	StockCode string
	ISBN      string
	Author    Author
}
type Author struct {
	ID   string
	Name string
}

func remove(s []books, i int) []books {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
