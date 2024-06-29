package main

import "fmt"

func main() {

  intSlice := [] int{0,1,2,3,4,5,6,7,8,9}

  for _, num := range intSlice {
    if num % 2 == 0 {
      //even
      fmt.Printf("Number %d is even", num) 
    } else {
      fmt.Printf("Number %d is odd", num)
    }
  }
}
