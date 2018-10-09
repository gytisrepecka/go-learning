// Copyright 2018 Gytis Repečka. All rights reserved.
// Use of this source code is governed by a GNU GPL
// license that can be found in the LICENSE file.

package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return n*factorial(n-1)
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Input number (and press ENTER): ")
  inputString, _ := reader.ReadString('\n')

  // Convert CRLF to LF
  // Windows
  // inputString = strings.Replace(inputString, "\r\n", "", -1)
  // Unix
  inputString = strings.Replace(inputString, "\n", "", -1)

  // Convert string (inputString) to integer (inputInt)
  var inputInt int
  inputInt, err := strconv.Atoi(inputString)

  // If there was error converting string to int
  if err != nil {
		fmt.Println(err)
	}

  fmt.Printf("Value (string): %s\n", inputString)
  fmt.Printf("Value length: %d\n", len(inputString))
  fmt.Printf("Value (int): %d\n", inputInt)
  fmt.Println("\n---------------")

  fmt.Printf("Factorial of %d is: %d", inputInt, factorial(inputInt))
  fmt.Println("\n---------------\n")
}
