package main

import (
	"fmt"
	"github.com/ceub/avaliacao-1/bonus"
	"github.com/ceub/avaliacao-1/q1"
	"github.com/ceub/avaliacao-1/q2"
	"github.com/ceub/avaliacao-1/q3"
	"github.com/ceub/avaliacao-1/q4"
	"github.com/ceub/avaliacao-1/q5"
)

func main() {
	possible, err := q1.DivideWatermelon(100)
	fmt.Printf("Q1:\tpossible: %v,\terr: %v\n", possible, err)

	solved := q2.ProblemsSolved([][3]bool{{true, true, true}, {false, false, false}, {true, false, true}})
	fmt.Printf("Q2:\tsolved: %v\n", solved)

	pieces, err := q3.DominoPieces(5, 5)
	fmt.Printf("Q3:\tpieces: %v,\terr: %v\n", pieces, err)

	order, err := q4.ClassifyPrices([]int{5, 6, 10})
	fmt.Printf("Q4:\torder: %v,\terr: %v\n", order, err)

	processString := q5.ProcessString("HelloWorld")
	fmt.Printf("Q5:\tprocessString: %v\n", processString)

	maxHeight, towers := bonus.CalculateTowers([]int{6, 5, 6, 7})
	fmt.Printf("Bonus:\tmaxHeight: %v,\ttowers: %v", maxHeight, towers)
}
