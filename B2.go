/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-19
 * @fileoverview Buying a car with extra additives.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// declare variables
	var basePrice float64 = 25000.00

	const floorMat float64 = 500.00
	var choiceFloor string

	const navSystem float64 = 1000.00
	var choiceNav string

	const leatherSeats float64 = 500.00
	var choiceLeather string

	const WARRANTY float64 = 2500.00
	var choiceWarranty string

	reader := bufio.NewReader(os.Stdin)

	// input

	// floor mats
	fmt.Print("Would you like to add floor mats to your car? \n")
	choiceFloor, _ = reader.ReadString('\n')
	choiceFloor = strings.TrimSpace(choiceFloor)

	// navigation system
	fmt.Print("Would you like to add a navigation system to your car? \n")
	choiceNav, _ = reader.ReadString('\n')
	choiceNav = strings.TrimSpace(choiceNav)

	// leather seats
	fmt.Print("Would you like to add heated leather seats to your car? \n")
	choiceLeather, _ = reader.ReadString('\n')
	choiceLeather = strings.TrimSpace(choiceLeather)

	// warranty
	fmt.Print("Would you like to add an extended 5-year warranty to your car? \n")
	choiceWarranty, _ = reader.ReadString('\n')
	choiceWarranty = strings.TrimSpace(choiceWarranty)

	// output

	// base price
	fmt.Printf("Base price: %.2f \n", basePrice)

	// floor mats
	if strings.ToLower(choiceFloor) == "yes" {
		basePrice = basePrice + floorMat
		fmt.Printf("Floor mats: %.2f \n", floorMat)
	} else {
	}

	// navigation system
	if strings.ToLower(choiceNav) == "yes" {
		basePrice = basePrice + navSystem
		fmt.Printf("Navigation system: %.2f \n", navSystem)
	} else {
	}

	// leather seats
	if strings.ToLower(choiceLeather) == "yes" {
		basePrice = basePrice + leatherSeats
		fmt.Printf("Heated leather seats: %.2f \n", leatherSeats)
	} else {
	}

	// warranty
	if strings.ToLower(choiceWarranty) == "yes" {
		basePrice = basePrice + WARRANTY
		fmt.Printf("5-year extended warranty: %.2f \n", WARRANTY)
	} else {
	}

	// tax
	var tax float64 = basePrice * 0.13
	fmt.Printf("Tax: %.2f \n", tax)

	// total price
	var totalPrice float64 = basePrice + tax
	fmt.Printf("Total: %.2f \n", totalPrice)

	fmt.Println("\nDone.")
}
