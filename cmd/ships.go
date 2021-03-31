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
	var loopTrack int = 0
	// make random orientation rows or columns
	//rows
	rand.Seed(time.Now().UnixNano())
	if true {
		randRow := rand.Intn(8)
		randCol := rand.Intn(8)
		if randRow+ship.size <= 8 {
			fmt.Println("it fits row", randRow, randCol)
			for randRowTest := randRow; loopTrack < ship.size; loopTrack++ {
				fmt.Println("test", randRowTest)
				randRowTest = randRowTest + 1
				if grid.Rows[randRowTest].Cells[randCol].Ship { //check if ship is already on it
					PlaceShips(grid, ship)
					fmt.Println("ship already on row")
					break
				}
			}

			//so first I see if there are any ships on the squares. If so then restart and break the loop
			// if no true is found then you can proceed to place the ships
			loopTrack = 0 //needed to reset looptrack
			for randRowPlace := randRow; loopTrack < ship.size; loopTrack++ {
				fmt.Println("placing", randRowPlace)
				randRowPlace = randRowPlace + 1
				grid.Rows[randRowPlace].Cells[randCol].Ship = true
				grid.Rows[randRowPlace].Cells[randCol].ShipName = ship.name
				fmt.Println("placing ship on", randRowPlace, randCol)
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

			for randColTest := randCol; loopTrack < ship.size; loopTrack++ {
				randColTest = randColTest + 1
				if grid.Rows[randRow].Cells[randColTest].Ship {
					fmt.Println("ship already on col")
					PlaceShips(grid, ship)
					break
				}
			}

			loopTrack = 0 //needed to reset looptrack

			for randColPlace := randCol; loopTrack < ship.size; loopTrack++ {
				randColPlace = randColPlace + 1
				grid.Rows[randRow].Cells[randColPlace].Ship = true
				grid.Rows[randRow].Cells[randColPlace].ShipName = ship.name
				fmt.Println(grid.Rows[randRow].Cells[randColPlace].ShipName)

			}
		} else {
			fmt.Println("it doesn't fit")
			PlaceShips(grid, ship)
		}
	}

	newGrid = grid

	return newGrid

}
