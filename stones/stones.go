package stones

import (
	"fmt"
	"log"
)

func GetJewelry() {

	var s, j string

	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		log.Fatalf("error on scan string: %s\n", err)
	}
	_, err = fmt.Scanf("%s", &j)
	if err != nil {
		log.Fatalf("error on scan string: %s\n", err)
	}

	seen := map[rune]struct{}{}
	for _, letter := range s {
		seen[letter] = struct{}{}
	}

	result := 0
	for _, stone := range j {
		if _, ok := seen[stone]; ok {
			result++
		}
	}

	fmt.Println(result)
}
