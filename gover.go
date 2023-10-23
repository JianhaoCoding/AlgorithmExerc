package main

import {
  "fmt"
	"strconv"
}

func main() {
    details := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
	  fmt.Println(countSeniors(details))
}

// 老人的数量
func countSeniors(details []string, age int) int {
    count := 0
    for _, detail = range details {
      nowAge, err = strconv.Atoi(detail[11, 13])
      if err == nil && nowAge >= age {
          count++
      }
    }
    return count
}
