package main

var theCells [9]Cell
var rowToAdd Row
var theRows [9]Row
var theGrid Grid

type Cell struct {
	RowNumber int
	ColNumber int
	Ship      bool
	ShipName  string
}

type Row struct {
	Cells [9]Cell
}

type Grid struct {
	Rows [9]Row
}

func CreateGrid() Grid {

	for i := 0; i < 9; i++ {
		for y := 0; y < 9; y++ {
			var newCell = createSquare(i, y)
			theCells[y] = newCell
			rowToAdd = Row{theCells}
		}
		theRows[i] = rowToAdd
		theGrid = Grid{theRows}
	}
	return theGrid
}

func createSquare(colNumber, rowNumber int) Cell {

	var newCell = Cell{colNumber, rowNumber, false, "water"}
	return newCell

}
