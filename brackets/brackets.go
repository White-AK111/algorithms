package brackets

import (
	"bufio"
	"os"
	"strconv"
)

func DoBrackets() {
	n := 0

	//fr, err := os.Open("brackets/input.txt")
	fr, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	sc := bufio.NewScanner(fr)

	for sc.Scan() {
		iv, _ := strconv.Atoi(sc.Text())
		n = iv
	}

	//fw, err := os.Create("brackets/output.txt")
	fw, err := os.Create("output.txt")
	if err != nil {
		os.Exit(1)
	}
	defer fw.Close()

	recBracket(n, "", 0, 0, fw)
}

func recBracket(n int, res string, left int, right int, fw *os.File) {
	if left == n && right == n {
		fw.WriteString(res + "\r\n")
	} else {
		if left < n {
			recBracket(n, res+"(", left+1, right, fw)
		}
		if right < left {
			recBracket(n, res+")", left, right+1, fw)
		}
	}
}
