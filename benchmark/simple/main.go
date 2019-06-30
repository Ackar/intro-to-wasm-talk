package main

import (
  "time"
  "fmt"
)

func compute(n int) {
  var res int
  for i := 0; i < n; i++ {
    res += 2
  }
}

func main() {
  start := time.Now()
  compute(10000000000)
  fmt.Printf("Time: %s\n", time.Since(start))
}
