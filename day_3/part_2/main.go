package main

import (
  "fmt"
  "strconv"
  "os"
  "strings"
  "math/bits"
  "math"
)

func main() {
  input, err := getInput("../input.txt")
  checkError(err)

  processedInput, err := processInput(input)
  checkError(err)

  ogr := oxygenGeneratorRating(processedInput)
  csr := co2ScrubberRating(processedInput)
  fmt.Println(ogr[0]*csr[0])
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

func oxygenGeneratorRating(input []int64) []int64 {
  findMostCommonBit := func (zero, one int) int {
    if one > zero || one == zero {
      return 1
    } else {
      return 0
    }
  }
  return findRate(input, findMostCommonBit)
}

func co2ScrubberRating(input []int64) []int64 {
  findLeastCommon := func (zero, one int) int {
    if zero < one || one == zero {
      return 0
    } else {
      return 1
    }
  }
  return findRate(input, findLeastCommon)
}
 
func findRate(input []int64, f func(zero, one int)(int)) []int64 {
  nb := numberOfBits(input) - 1
  mask := int64(math.Exp2(float64(nb)))
  resp := input

  for bit:=nb; bit >= 0; bit-- {
    zero, one := countBits(resp, mask)
    bitFilter := f(zero, one)
    resp = filterValues(resp, bitFilter, mask)
    mask /= 2
  }
  return resp
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

func countBits(input []int64, mask int64) (int, int) {
  zero := 0
  one := 0

  for _, v := range input {
    if (v & mask) == mask {
      one++
    } else {
      zero++
    }
  }
  return zero, one
}

func filterValues(input []int64, bitFilter int, mask int64) []int64 {
  if len(input) == 1 {
    return input
  }

  var resp []int64

  for _, v := range input {
    if bitFilter == 1 && (v & mask) == mask {
      resp = append(resp, v)
    } else if bitFilter == 0 && (v & mask) != mask {
      resp = append(resp, v)
    }
  }
  return resp
}

func bitArrayToInt(a []int) (int64, error) {
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
