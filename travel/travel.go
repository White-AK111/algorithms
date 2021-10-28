package travel

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type City struct {
	x int
	y int
}

func FindWay() {
	mapCity := map[int]City{}
	countCity := 0
	currStr := 0
	maxCost := 0
	startCity := 0
	finishCity := 0

	fr, err := os.Open("travel/input.txt")

	if err != nil {
		panic(err)
	}
	defer fr.Close()

	sc := bufio.NewScanner(fr)
	const maxCapacity = 1024
	buf := make([]byte, maxCapacity)
	sc.Buffer(buf, maxCapacity)

	for sc.Scan() {
		currStr++
		iv, _ := strconv.Atoi(sc.Text())
		if currStr == 1 {
			countCity = iv
		}
		if currStr > 1 && currStr <= countCity+1 {
			str := strings.Split(sc.Text(), " ")
			x, _ := strconv.Atoi(str[0])
			y, _ := strconv.Atoi(str[1])
			currCity := City{x: x, y: y}
			mapCity[currStr-1] = currCity
		}
		if currStr == countCity+2 {
			maxCost, _ = strconv.Atoi(sc.Text())
		}
		if currStr == countCity+3 {
			str := strings.Split(sc.Text(), " ")
			startCity, _ = strconv.Atoi(str[0])
			finishCity, _ = strconv.Atoi(str[1])
		}
	}

	fmt.Printf("City count: %d\n", countCity)
	fmt.Printf("City: %v\n", mapCity)
	fmt.Printf("Cost: %d\n", maxCost)
	fmt.Printf("Start city: %d, Finish city: %d\n", startCity, finishCity)

	visitedCity := map[int]City{}
	visitedCity[startCity] = mapCity[startCity]
	lowCost := findLowestCostWay(mapCity, visitedCity, startCity, finishCity, maxCost)
	fmt.Printf("Lowest way: %d\n:", lowCost)

	for i := 0; i < len(visitedCity); i++ {
		fmt.Print(visitedCity[i])
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

func findLowestCostWay(mapCity map[int]City, visitedCity map[int]City, startCity int, finishCity int, maxCost int) int {

	for i := 1; i <= len(mapCity); i++ {
		if _, ok := visitedCity[i]; !ok {
			visitedCity[i] = mapCity[i]
			if _, ok := visitedCity[finishCity]; ok {
				return 0
			}
			return findLowestCostWay(mapCity, visitedCity, startCity, finishCity, maxCost)
		}
	}

	return 0
}
