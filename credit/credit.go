package credit

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetCredit() {
	res := 0
	var strArr1, strArr2 []string

	fr, err := os.Open("credit/input.txt")
	if err != nil {
		fmt.Println(strconv.Itoa(res))
		return
	}
	defer fr.Close()

	sc := bufio.NewScanner(fr)
	sc.Split(bufio.ScanLines)
	//const maxCapacity = 1024
	//buf := make([]byte, maxCapacity)
	//sc.Buffer(buf, maxCapacity)

	var txtLines []string
	for sc.Scan() {
		txtLines = append(txtLines, sc.Text())
	}

	strArr1 = strings.Split(txtLines[0], " ")
	strArr2 = strings.Split(txtLines[1], " ")

	notifyDay, _ := strconv.Atoi(strArr1[1])
	currNotify := 0
	maxNotify, _ := strconv.Atoi(strArr1[2])
	mapNotify := make(map[int]int)
	currN := 0
	var days []int

	sort.Strings(strArr2)

	for _, task := range strArr2 {
		currNotify, _ = strconv.Atoi(task)
		if _, ok := mapNotify[currNotify]; !ok {
			currN++
			mapNotify[currNotify] = currN
		}
		for j := 2; j <= maxNotify; j++ {
			currNotify = currNotify + notifyDay
			if _, ok := mapNotify[currNotify]; !ok {
				currN++
				mapNotify[currNotify] = currN
			}
		}
	}

	for k, _ := range mapNotify {
		days = append(days, k)
	}

	sort.Ints(days)
	res = days[maxNotify-1]
	fmt.Println(strconv.Itoa(res))
}
