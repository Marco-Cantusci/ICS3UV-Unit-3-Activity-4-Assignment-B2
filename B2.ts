/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-19
 * @fileoverview Buying a car with extra additives.
 */

// declare variables
let basePrice: number = 25000.00;

const floorMat: number = 500.00;

const navSystem: number = 1000.00;

const leatherSeats: number = 500.00;

const WARRANTY: number = 2500.00;

// input

// floor mats
const choiceFloor: string = prompt("Would you like to add floor mats to your car? \n") || ("\n");

// navigation system
const choiceNav: string = prompt("Would you like to add a navigation system to your car? \n") || ("\n");

// leather seats
const choiceLeather: string = prompt("Would you like to add heated leather seats to your car? \n") || ("\n");

// warranty
const choiceWarranty: string = prompt("Would you like to add an extended 5-year warranty to your car? \n") || ("\n");


// output

// base price
console.log(`Base price: ${basePrice.toFixed(2)}`);

// floor mats
if (choiceFloor.toLowerCase() == "yes") {
  basePrice = basePrice + floorMat;
  console.log(`Floor mats: ${floorMat.toFixed(2)}`);
} else {
}

// navigation system
if (choiceNav.toLowerCase() == "yes") {
  basePrice = basePrice + navSystem;
  console.log(`Navigation system: ${navSystem.toFixed(2)}`);
} else {
}

// leather seats
if (choiceLeather.toLowerCase() == "yes") {
  basePrice = basePrice + leatherSeats;
  console.log(`Heated leather seats: ${leatherSeats.toFixed(2)}`);
} else {
}

// warranty
if (choiceWarranty.toLowerCase() == "yes") {
  basePrice = basePrice + WARRANTY;
  console.log(`5-year extended warranty: ${WARRANTY.toFixed(2)}`);
} else {
}

// tax
const tax: number = basePrice * 0.13;
console.log(`Tax: ${tax.toFixed(2)}`);

// total price
const totalPrice: number = basePrice + tax;
console.log(`Total: ${totalPrice.toFixed(2)}`);

console.log("\nDone.");
