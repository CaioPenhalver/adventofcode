package main

import (
  "os"
  "fmt"
  "strings"
  "strconv"
)

// day 1 https://adventofcode.com/2021/day/1
// var input []int = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func main() {
  input, err := getInput()
  checkError(err)

  processedInput, err := processInput(input)
  checkError(err)

  fmt.Println(countIncreases(processedInput))
}

func countIncreases(input []int) int {
  counter := 0
  previousValue := input[0]

  for _, currentValue := range input {
    if currentValue > previousValue {
      counter++
  }
  previousValue = currentValue
  }
  return counter
}

func getInput() (string, error) {
  input, err := os.ReadFile("./input.txt")
  if err != nil {
    return "", err
  }

  return string(input), nil
}

func processInput(input string) ([]int, error) {
  var processedInput []int

  for _, v := range strings.Split(input, "\n") {
    if v != "" {
      i, err := strconv.Atoi(v)
      if err != nil {
        return nil, err
      }
      processedInput = append(processedInput, i)
    }
  }
  return processedInput, nil
}

func checkError(err error) {
  if err != nil {
    fmt.Printf("An error has ocurred: %v\n", err)
    os.Exit(1)
  }
}
