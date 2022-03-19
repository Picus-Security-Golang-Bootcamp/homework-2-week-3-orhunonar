## Homework | Week 3

The Struct of this project has
```
- Book ID
- Book Name
- Page Number
- Amount Book on the Stock
- Price
- Stock Code
- ISBN
- Information about the Author (ID and Name)
```

1. List Command : List the book
2. Search: Search the Book
3. Get : Get the information of book that has this ID
4. Delete: Delete the book that has this ID
5. Buy : Buy from the book with the specified id as much as the number written



### list command
```
go run main.go list
```

### search command 
```
go run main.go search <bookName>
go run main.go search Lord of the Ring: The Return of the King
```

### get command
```
go run main.go get <bookID>
go run main.go get 5
```

### delete command
```
go run main.go delete <bookID>
go run main.go delete 5
```

### buy command
```
go run main.go buy <bookID> <quantity>
go run main.go buy 5 2
```

