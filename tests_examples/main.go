
package main
import "fmt"

func Sum(data []int) int {
  s := 0
  for _, val := range data {
    s += val
  }
  return s
}


func main() {
  data := []int{ 1, 2, 4, -7}
  fmt.Println(Sum(data))
}
