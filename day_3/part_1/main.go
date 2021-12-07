package main

import (
  "fmt"
  "strconv"
  "os"
  "strings"
  "math/bits"
)

func main() {
  input, err := getInput("../input.txt")
  checkError(err)

  processedInput, err := processInput(input)
  checkError(err)

  ones := countOnes(processedInput)
  threshold := float64(len(processedInput)) / 2

  gr := findGammaRate(ones, threshold)
  er := findEpsilonRate(ones, threshold)
  fmt.Println(gr)
  fmt.Println(er)

  m, err := multiply(gr, er)
  checkError(err)

  fmt.Println(m)
}

func getInput(path string) (string, error) {
  input, err := os.ReadFile(path)
  if err != nil {
    return "", err
  }
  return string(input), nil
}

func processInput(input string) ([]int64, error) {
  inputArr := strings.Split(input, "\n")
  inputArr = deleteInvalidValues(inputArr)
  processedArr, err := stringBitsToInt(inputArr)
  if err != nil {
    return nil, err
  }
  return processedArr, nil
}

func deleteInvalidValues(input []string) (processedInput []string) {
  for _, v := range input {
    if v != "" {
      processedInput = append(processedInput, v)
    }
  }
  return
}

func stringBitsToInt(input []string) ([]int64, error) {
  var intArr []int64
  for _, v := range input {
    i, err := strconv.ParseInt(v, 2, 64)
    if err != nil {
      return nil, err
    }
    intArr = append(intArr, i)
  }
  return intArr, nil
}


func countOnes(input []int64) []int {
  nb := numberOfBits(input)
  counter := make([]int, nb, nb)
  mask := int64(1)

  for bit:=nb-1; bit >= 0; bit-- {
    for _, v := range input {
      if (v & mask) == mask {
        counter[bit]++
      }
    }
    mask *= 2
  }
  return counter
}

func numberOfBits(input []int64) int {
  gv := int64(0)
  for _, v := range input {
    if v > gv {
      gv = v
    }
  }
  return bits.Len64(uint64(gv))
}

func findGammaRate(ones []int, threshold float64) []int {
  var gr []int

  for _, v := range ones {
    if float64(v) > threshold {
      gr = append(gr, 1)
    } else {
      gr = append(gr, 0)
    }
  }
  return gr
}

func findEpsilonRate(ones []int, threshold float64) []int {
  var er []int

  for _, v := range ones {
    if float64(v) < threshold  && v > 0 {
      er = append(er, 1)
    } else {
      er = append(er, 0)
    }
  }
  return er

} 

func multiply(gr , er []int) (int64, error) {
  gri, err := arrayByteToInt(gr)
  if err != nil {
    return 0, err
  }

  eri, err :=arrayByteToInt(er)
  if err != nil {
    return 0, err
  }

  return (gri * eri), nil
}

func arrayByteToInt(a []int) (int64, error) {
  strByte := strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), ""), "[]")
  i, err := strconv.ParseInt(strByte, 2, 64)
  if err != nil {
    return 0, err
  }
  return i, nil
}

func checkError(err error) {
  if err != nil {
    fmt.Printf("An error has ocurred: %v\n", err)
    os.Exit(1)
  }
}
