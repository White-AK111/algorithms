package main

import (
	"github.com/White-AK111/algorithms/anagrams"
	"github.com/White-AK111/algorithms/brackets"
	"github.com/White-AK111/algorithms/duplicates"
	"github.com/White-AK111/algorithms/stones"
	"github.com/White-AK111/algorithms/units"
)

func main() {
	stones.GetJewelry()
	units.GetMaxUnits()
	duplicates.DeleteDuplicates()
	brackets.DoBrackets()
	anagrams.CheckAnagrams()
}
