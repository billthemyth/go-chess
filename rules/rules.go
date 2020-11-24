package rules

import "log"

func clearPosition(column int, row  int, table *Table){
	table.Positions[column][row].SetFilledBy(Part{})
	table.Positions[column][row].SetIsFilled(false)
}

func checkIfPositionIsFilled(column int, row int, table *Table){
	if !table.Positions[column][row].isFilled {
		 log.Fatal("Position is not filled")
	}
}



