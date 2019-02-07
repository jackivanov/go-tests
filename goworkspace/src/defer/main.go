package main

import "fmt"

func main(){
  defer three()
  defer two()
  one()
}


func one(){
  fmt.Println("ONE")
}

func two(){
  fmt.Println("TWO")
}

func three(){
  fmt.Println("THREE")
}
