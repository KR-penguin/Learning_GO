package myfunctions

import "fmt"

func Swap(a, b *int) {

  defer fmt.Println("Swap Done ;)")

  var temp int = 0

  temp = *a
  *a = *b
  *b = temp
  
}