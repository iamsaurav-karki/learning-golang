package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Slice examples")

	var fruitLists = []string{"Apple", "mango", "peach", "peer"}
	fmt.Printf("type is %T \n", fruitLists)
	fmt.Println(fruitLists)

	fruitLists = append(fruitLists, "pineapple")
	fmt.Println("After append", fruitLists)

	// fruitLists = append(fruitLists[1:])
	fruitLists = append(fruitLists[:3])

	fmt.Println(fruitLists)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 236
	highScores[2] = 2334
	highScores[3] = 945

	highScores = append(highScores, 555, 434, 2312) // reallocates the memory so it can accomodate more than the initial 4.
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
