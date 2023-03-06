package helpers

import (
  "os"
  "strconv"
  "strings"
)

func Min(x, y int) int {
 if x < y {
   return x
 }
 return y
}

func Max(x, y int) int {
 if x > y {
   return x
 }
 return y
}

func GetEnv(key string, fallback int) int {
  value, ok := os.LookupEnv(key)
  if ok {
    parsedValue, _ := strconv.Atoi(value)
    return parsedValue
  }
  return fallback
}

func ParseUserInput(input string) (int, int) {
  split := strings.Split(input, ",")
  x, _ := strconv.Atoi(split[0])
  y, _ := strconv.Atoi(split[1])
  return x, y
}
