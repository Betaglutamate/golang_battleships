package main

import (
	"fmt"
)

func StartGame() {

	// Create the Grid
	theGrid := CreateGrid()
	// Place the Ships

	//ask for user name
	fmt.Println("Enter Your Name: ")

	// var then variable name then variable type
	var name string

	// Taking input from user
	fmt.Scanln(&name)

	//set ships
	for index, ship := range ShipList {
		fmt.Println(index, ship.name)
		theGrid = PlaceShips(theGrid, ship)
	}

	//take a shot
	rowNumber, colNumber := shoot()

	fmt.Println(theGrid.Rows[rowNumber].Cells[colNumber])

}

func shoot() (int, int) {
	fmt.Println("guess a row (1-9): ")
	var guessRow int
	fmt.Scanln(&guessRow)

	fmt.Println("guess a column (1-9): ")
	var guessColumn int
	fmt.Scanln(&guessColumn)
	return guessRow, guessColumn
}
