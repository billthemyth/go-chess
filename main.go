package main

import (
	rules "./rules"
)

func main() {


	table := rules.Table{
		Positions: [8][8]rules.Position{},
	}

	table.Positions[7][7] = rules.Position{}
	table.Positions[7][7].SetIsFilled(true)
	table.Positions[7][7].SetFilledBy(rules.Part{IsDead: false})

	rules.PrintTable(&table)




}