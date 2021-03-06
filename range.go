package main
import (
   "fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
   for i, v := range pow {
      fmt.Printf("2**%d = %d\n", i, v)
   }
   temp := make([]int, 10)
   for i := range temp {
      temp[i] = 1 << uint(i)
   }

   for _, value := range temp {
      fmt.Printf("%d\n", value)
   }
}

func printSlice(s string, x []int){
   fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}