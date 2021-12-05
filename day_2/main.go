package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main(){
  input, err := getInput("./input.txt")
  checkError(err)

  commands, err := processInput(input)
  checkError(err)

  s := &Submarine {
    Horizontal: 0,
    Depth:      0,
  }
  s.runCommands(commands)

  fmt.Println(s.Horizontal * s.Depth)
}

type Command struct {
  Name  string
  Value int
}

type Submarine struct {
  Horizontal int
  Depth      int
}

func (s *Submarine) runCommands(commands []*Command) {
  for _, command := range commands {
    switch command.Name {
    case "forward":
      s.Horizontal += command.Value
    case "up":
      s.Depth -= command.Value
    case "down":
      s.Depth += command.Value
    }
  }
}

func getInput(path string) (string, error) {
  input, err := os.ReadFile(path)
  if err != nil {
    return "", err
  }

  return string(input), nil
}

func processInput(input string) ([]*Command, error) {
  inputArr := strings.Split(input, "\n")
  inputArr = deleteInvalidValues(inputArr)

  return parseCommands(inputArr)
}

func deleteInvalidValues(input []string) (processedInput []string) {
  for _, v := range input {
    if v != "" {
      processedInput = append(processedInput, v)
    }
  }
  return
}

func parseCommands(input []string) ([]*Command, error) {
  var commands []*Command

  for _, command := range input {
    parsedCommand, err := parseCommand(command)
    if err != nil {
      return nil, err
    }
    commands = append(commands, parsedCommand)
  }

  return commands, nil
}

func parseCommand(command string) (*Command, error) {
  commandValue := strings.Split(command, " ")
  valueInt, err := strconv.Atoi(commandValue[1])
  if err != nil {
    return nil, err
  }
  return &Command{commandValue[0], valueInt}, nil
}

func checkError(err error) {
  if err != nil {
    fmt.Printf("An error has ocurred: %v\n", err)
    os.Exit(1)
  }
}
