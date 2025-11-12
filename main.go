/* 
 * Author Natnael Debesay
 * Version 1.0.0
 * Date 2025-11-11
 * This program finds the area of a circle.
 */

package main

import "fmt"

func main() {
	// initialize numbers as constants
	const RADIUS float64 = 5.6
	const PI float64 = 3.14

	// INPUT - none

	// PROCESS 
	// calulate the area
	answer := RADIUS * RADIUS * PI

	// OUTPUT
	// display the result
	fmt.Println("The area of a circle with a radius of " , 5.6, "is:", answer,)

	fmt.Println("\nDone.")
}
