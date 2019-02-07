package main

import (
	"fmt"
)

type book struct {
	author string
	year   string
	pages  int
	price  int
}


func (b book) ChangePrice(){
  fmt.Println(b.price)
  b.price = b.price + 2
  fmt.Println(b.price)
}

func (b *book) ChangePages(){
  fmt.Println(b.pages)
  b.pages = b.pages + 500
  fmt.Println(b.pages)
}

func main() {
  FirstBook := book{
    author: "Mr. Smith",
    year: "2018",
    pages: 200,
    price: 18,
  }

  SecondBook := new(book)
  SecondBook.author = "123"
  SecondBook.year = "2019"
  SecondBook.pages = 100
  SecondBook.price =1

  // SecondBook := book{
  //   author: "Mr. Smith",
  //   year: "2018",
  //   pages: 100,
  //   price: 1,
  // }

  FirstBook.ChangePrice()

  SecondBook.ChangePages()

  SecondBook.ChangePages()
}
