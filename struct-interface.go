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

type Truck struct {
  modelName string
	year int
	started bool
  weight int
}

// Interfaces
/*
Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
Good example of interfaces: https://gobyexample.com/interfaces
*/
// Vehicle will act as interface to Car and Truck structs.
type Vehicle interface {
	FriendlyName() string
	Wheels() int
  Start()
}

// Functions (methods)

// Receiving a pointer to Car/Truck so can modify it (change started value)

// Start method for Car
func (c *Car) Start() {
	c.started = true
	fmt.Printf("Car %s %s says vrooom!\n", c.make.makeTitle, c.model)
}

// Start method for Truck
func (t *Truck) Start() {
	t.started = true
	fmt.Printf("Truck %s (max cargo: %d t) says br br br br!\n", t.modelName, t.weight)
}

// Because defined only for Car and not listed as method in
// type Vehicle interface, StopCar is applicable only for Car
// but not Vehicle.
func (c *Car) StopCar() {
	if c.started == true {
		c.started = false
		fmt.Printf("Car %s %s was stopped.\n", c.make.makeTitle, c.model)
	} else {
		fmt.Printf("Car %s %s is not running anyway.\n", c.make.makeTitle, c.model)
	}
}

// This method is currently implemented for Car only. There is no implementation
// for Truck therefore cannot use this method for Vehicle.
func (cr *Car) AlterName(nameAppend string) (string) {
  if nameAppend != "" {
    cr.model = cr.model + " " + nameAppend
  }
	outString := fmt.Sprintf("%s %s (%d)", cr.make.makeTitle, cr.model, cr.year)
	return outString
}

// Car has 4 wheels
func (cr Car) Wheels() int {
	return 4
}

// Truck has 6 wheels
func (tr Truck) Wheels() int {
	return 6
}
// ...and we can then use Wheels method to know how many wheels Vehicle has.

func (cr Car) FriendlyName() (string) {
	outStarted := "not running"
	if cr.started == true {
		outStarted = "vuu tuuuu tuuuuu"
	}
	outString := fmt.Sprintf("%s %s (%d) %s", cr.make.makeTitle, cr.model, cr.year, outStarted)
	return outString
}

func (t Truck) FriendlyName() (string) {
	outStarted := "not running"
	if t.started == true {
		outStarted = "vuu tuuuu tuuuuu"
	}
	outString := fmt.Sprintf("%s (%d) %s", t.modelName, t.year, outStarted)
	return outString
}

// This method is for Vehicle therefore universal (not specifically for Car or
// for Truck). Mind that inside we can use only those methods which are defined
// for Vehicle (i.e. for Car and for Truck).
func VehicleDetails(vh Vehicle) {
	fmt.Printf("Vehicle %s has %d wheels.\n", vh.FriendlyName(), vh.Wheels())
}

// This method is intended for interface Vehicle and works for both Car and Truck
func VehicleStart(vh Vehicle) {
  fmt.Println("Starting vehicle [from interface method]...")
  vh.Start()
}


// Main function (execution point)

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


	fmt.Println("Will try out interfaces below:")

  // Create new vehicle as existing car (Toyota Yaris (cars[2])
  /*
    Using address pointer (&) to cars[2] meaning that if we modify newVehicle
    values, cars[2] will be affected because newVehicle is not a copy of cars[2]
    but the same instance.
    If we write newVehicle := cars[2] then "Read from cars array" would output
    unchanged value (not "Yaris Limited Edition" but still "Yaris" even though
    in newVehicle we whould have "Yaris Limited Edition").

    More info about pointers: http://piotrzurek.net/2013/09/20/pointers-in-go.html
  */
  // newVehicle := Car{carMakes[0], "RAV4", 2018, false}
  // newVehicle := cars[2]
  newVehicle := &cars[2]


	fmt.Println("Starting newVehicle...")
  // Since newVehicle is Car, we can apply Car's methods (Start and AlterName)
	newVehicle.Start()
  newVehicle.AlterName("Limited Edition")

  // This method is defined for Vehicle therefore can be used on Car or Truck.
  // VehicleDetails receives a copy of object (does not require pointer) so as
  // underlying methods (FriendlyName and Wheels) that are called from VehicleDetails.
	VehicleDetails(newVehicle)
  fmt.Println("--------------------------")

  fmt.Println("Read from cars array:")
  fmt.Printf("%s %s\n", cars[2].make.makeTitle, cars[2].model)
  // Because pointer was used, Car.model was updated.

  fmt.Println("--------------------------")
  fmt.Println("Let's try out Vehicle interface and it's method:")
  fmt.Println("--------------------------")

  // Let's create two Vehicle instances:
  // - a Car using a pointer to Toyota Yaris
  vehicleCar := &cars[2]
  // - a new Truck
  vehicleTruck := Truck{"Mercedes Actros", 2018, false, 10}

  VehicleStart(vehicleCar)
  /*
    Must use pointer here - &vehicleTruck because Start expects pointer.
    VehicleCar is already a pointer, so can use VehicleStart(vehicleCar)
    whereas vehicleTruck we create now. If using VehicleStart(vehicleTruck)
    (without pointer) this error shown:
    .\struct-interface.go:229:15: cannot use vehicleTruck (type Truck) as type Vehicle in argument to VehicleStart:
        Truck does not implement Vehicle (Start method has pointer receiver)
  */
  VehicleStart(&vehicleTruck)
  fmt.Println("--------------------------")
  fmt.Println("Let's now show Vehicle in nicely formatted way:")

  // As explained above, VehicleDetails and underlying methods receives a copy
  // of object but we can pass a pointer to it as well.
  VehicleDetails(&vehicleTruck)

}
