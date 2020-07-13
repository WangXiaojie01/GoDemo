package main
import (
   "fmt"
)

type Vertex struct {
   X int
   Y int
}

func main() {
   fmt.Println(Vertex{1, 2})
   ver := Vertex{3, 4}
   fmt.Println(ver.X)
   fmt.Println(ver.Y)

   ver.X = 7
   fmt.Println(ver.X)
   fmt.Println(ver.Y)

   p := &ver 
   p.X = 1e9
   fmt.Println(ver.X)
   fmt.Println(ver.Y)
}
