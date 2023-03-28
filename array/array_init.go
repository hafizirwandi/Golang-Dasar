package main

import "fmt"

func main(){
  arr0 := []int{} //not initialized
  arr1 := [5]int{} //not initialized
  arr2 := [5]int{1,2} //partially initialized
  arr3 := [5]int{1,2,3,4,5} //fully initialized
  fmt.Println(arr0)
  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)


  
}

// output
//[]
// [0 0 0 0 0]
// [1 2 0 0 0]
// [1 2 3 4 5]
