package rules

import (
	"fmt"
	"strconv"
)

const COLUMNS = 8
const ROWS	= 8

type Table struct {
	Positions [COLUMNS][ROWS]Position
}

type Position struct {

	isFilled bool
	filledBy Part

}

func (p *Position) IsFilled() bool {
	return p.isFilled
}

func (p *Position) SetIsFilled(isFilled bool) {
	p.isFilled = isFilled
}

func (p *Position) FilledBy() Part {
	return p.filledBy
}

func (p *Position) SetFilledBy(filledBy Part) {
	p.filledBy = filledBy
}

func PrintTable(table *Table){

	for i := 0; i < ROWS ; i++ {

		fmt.Print("|")

		for j := 0; j < COLUMNS ; j++ {

			if table.Positions[j][i].isFilled {

				fmt.Print(strconv.Itoa(j) + "c-" + strconv.Itoa(i)+"r")

			} else {
				fmt.Print(" ")
			}

			fmt.Print("|")
		}
		fmt.Println()
	}
	fmt.Println()
}


