package main
import "fmt"

func main() {
  for i := 0; i < 3; i++ {
    println(i)
  }
  

  for i, t := range []int{5,6,7} {
    fmt.Printf("i=%d, t= %d\n", i, t)
  }
  
  w := 0
  for w < 2 {
    println(w)
    w++
  }
  
}
