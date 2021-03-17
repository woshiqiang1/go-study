package main

import (
    "fmt"
)

type Book struct {
    title string
    author string
  }
  
  func main() {
    var myBook Book
    myBook = Book{
      title: "《两京十五日》",
      author: "马伯庸",
    }
    printBook(myBook)
  }
  
  func printBook(book Book) {
    fmt.Printf("book title: %s\n", book.title)
    fmt.Printf("book author: %s\n", book.author)
  }
