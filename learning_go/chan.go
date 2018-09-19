package main
 
import "fmt"
 
func main() {
   count := 40
   fmt.Println("Final counter is", <-increment(<-increment(count)))
}
 
func increment(n int) chan int {
   out := make(chan int)
   go func(count int) {
      for i := 0; i < 20; i++ {
         count++
      }
      out <- count
   }(n)
 
   return out
}
