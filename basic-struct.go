// Copyright 2018 Gytis Repečka. All rights reserved.
// Use of this source code is governed by a GNU GPL
// license that can be found in the LICENSE file.

package main

import (
  "fmt"
)

// Define named structures

type CarMake struct {
  makeTitle string
  countryOfOrigin string
}

type Car struct {
	make CarMake
  model string
	year int
	started bool
}

// Functions (methods)

// Receiving a pointer to Car so can modify it
func (c *Car) Start() {
	c.started = true
	fmt.Printf("Car %s %s says vrooom!\n", c.make.makeTitle, c.model)
}

func main() {
  var carMakes []CarMake
  newCarMake := CarMake{makeTitle: "Toyota", countryOfOrigin: "Japan"}
  carMakes = append(carMakes, newCarMake)
  newCarMake = CarMake{makeTitle: "Renault", countryOfOrigin: "France"}
  carMakes = append(carMakes, newCarMake)
  newCarMake = CarMake{"Peugeot", "France"}
  carMakes = append(carMakes, newCarMake)

  var cars []Car
  newCar := Car{carMakes[1], "Clio", 1994, false}
  cars = append(cars, newCar)
  newCar = Car{carMakes[2], "406", 1997, false}
  cars = append(cars, newCar)
  newCar = Car{carMakes[0], "Yaris", 2000, false}
  cars = append(cars, newCar)

  fmt.Println("--------------------------")
  fmt.Println("Make | Model | Country")
  fmt.Println("--------------------------")

  for i:=0; i<len(cars); i++ {
    fmt.Printf("%s | %s | %s\n", cars[i].make.makeTitle, cars[i].model, cars[i].make.countryOfOrigin)
  }
  fmt.Println("--------------------------")

  // Let's execute struct function (method)
  cars[2].Start()
  cars[0].Start()
  fmt.Println("--------------------------")

  fmt.Println("Started cars:")
  startedCars := 0
  for i:=0; i<len(cars); i++ {
    if cars[i].started == true {
      fmt.Printf("%s %s\n", cars[i].make.makeTitle, cars[i].model)
      startedCars++
    }
  }
  if startedCars == 0 {
    fmt.Println("No cars are started.")
  } else {
    fmt.Printf("Total: %d\n", startedCars)
  }
  fmt.Println("--------------------------")
}
