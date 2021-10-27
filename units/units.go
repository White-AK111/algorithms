package units

import (
	"bufio"
	"os"
	"strconv"
)

func GetMaxUnits() {
	maxN := 0
	currN := 0
	n := 0

	//fr, err := os.Open("units/input.txt")
	fr, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	sc := bufio.NewScanner(fr)
	for sc.Scan() {
		if n == 0 {
			n++
		} else {
			currS, _ := strconv.Atoi(sc.Text())
			if currS == 1 {
				currN++
			} else {
				currN = 0
			}
			if currN >= maxN {
				maxN = currN
			}
		}
	}

	fw, err := os.Create("output.txt")

	if err != nil {
		os.Exit(1)
	}
	defer fw.Close()
	fw.WriteString(strconv.Itoa(maxN))
}
