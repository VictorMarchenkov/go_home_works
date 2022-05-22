// Package app developed for education purposes only.
// The example as simple as possible, but expose all function needed
//
// Below you can find usage example
package app

import (
	"fmt"
	"gb_home/home2/internal/controller"
)

// Run the place where all dirty work happen
func Run(a, b int) {
	fmt.Println("The main logic heare")
	res := controller.Add(a, b)
	fmt.Printf("%d + %d = %d", a, b, res)
}
