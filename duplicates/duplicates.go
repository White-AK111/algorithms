package duplicates

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func DeleteDuplicates() {
	n := 0
	m := map[int]struct{}{}

	//fr, err := os.Open("duplicates/input.txt")
	fr, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer fr.Close()

	sc := bufio.NewScanner(fr)
	const maxCapacity = 1024
	buf := make([]byte, maxCapacity)
	sc.Buffer(buf, maxCapacity)

	for sc.Scan() {
		iv, _ := strconv.Atoi(sc.Text())
		if n == 0 {
			n = iv
		} else {
			if _, ok := m[iv]; !ok {
				m[iv] = struct{}{}
			}
		}
	}

	res := make([]int, 0, len(m))
	for k, _ := range m {
		res = append(res, k)
	}
	sort.Ints(res)

	//fw, err := os.Create("duplicates/output.txt")
	fw, err := os.Create("output.txt")

	if err != nil {
		os.Exit(1)
	}
	defer fw.Close()
	for i := 0; i < len(res); i++ {
		fw.WriteString(strconv.Itoa(res[i]) + "\r\n")
	}
}
