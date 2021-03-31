package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ship struct {
	name string
	size int
}

var Destroyer = Ship{"Destroyer", 5}
var Cruiser = Ship{"Cruiser", 4}
var Submarine = Ship{"Submarine", 3}
var Small = Ship{"Small", 2}
var Patrol = Ship{"Patrol", 1}

var ShipList = [...]Ship{Destroyer, Cruiser, Submarine, Small, Patrol}

func rand1() bool {
	return rand.Float32() < 0.5
}

func PlaceShips(grid Grid, ship Ship) (newGrid Grid) {

	//get ship

	// make random orientation rows or columns
	//rows
	rand.Seed(time.Now().UnixNano())
	if rand1() {
		randRow := rand.Intn(8)
		randCol := rand.Intn(8)
		if randRow+ship.size <= 8 {
			fmt.Println("it fits row", randRow, randCol)
			for randRow := randRow; randRow < ship.size; randRow++ {
				if grid.Rows[randRow].Cells[randCol].Ship { //check if ship is already on it
					PlaceShips(grid, ship)
					break
				}
			}

			//so first I see if there are any ships on the squares. If so then restart and break the loop
			// if no true is found then you can proceed to place the ships
			for randRow := randRow; randRow < ship.size; randRow++ {
				grid.Rows[randRow].Cells[randCol].Ship = true
				grid.Rows[randRow].Cells[randCol].ShipName = ship.name
			}

		} else {
			fmt.Println("it doesn't fit")
			PlaceShips(grid, ship)
		}

	} else {

		randRow := rand.Intn(8)
		randCol := rand.Intn(8)
		if randCol+ship.size <= 8 {
			fmt.Println("it fits col", randRow, randCol)

			for randCol := randCol; randCol < ship.size; randCol++ {
				if grid.Rows[randRow].Cells[randCol].Ship {
					PlaceShips(grid, ship)
					break
				}
			}

			for randCol := randCol; randCol < ship.size; randCol++ {
				grid.Rows[randRow].Cells[randCol].Ship = true
				grid.Rows[randRow].Cells[randCol].ShipName = ship.name
			}
		} else {
			fmt.Println("it doesn't fit")
			PlaceShips(grid, ship)
		}
	}

	newGrid = grid

	fmt.Println(newGrid)

	return newGrid

	// find space on the grid for it make a random number 0-8

	//check if

	// grid.Rows[1].Cells[1].Ship = true
}
