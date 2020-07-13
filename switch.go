package main
import (
   "fmt"
   "runtime"
   "time"
)


func main() {
   fmt.Println("Go runs on")
   switch os := runtime.GOOS; os {
   case "darwin":
      fmt.Println("%s.", os)
      fmt.Println("OS X.")
   case "linux":
      fmt.Println("Linux")
   default:
      fmt.Println("%s", os)
   }

   t := time.Now()
   switch {
   case t.Hour() < 12:
      fmt.Println("Good morning")
   case t.Hour() < 17:
      fmt.Println("Good afternoon")
   default:
      fmt.Println("Good evening")
   }
}
