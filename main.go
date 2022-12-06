package main

import (
	"III/mve"
	"math/rand"
)

func main() {
	//print 50 in bold and red color
	println("\033[1m\033[31m", "Size 50", "\033[0m")
	A := populateAndShuffleArray(50)
	println(mve.BruteForceMaximizeExpression(A))
	println(mve.DynamicProgrammingMaximizeExpression(A))
	println()

	println("\033[1m\033[31m", "Size 100", "\033[0m")
	A = populateAndShuffleArray(100)
	println(mve.BruteForceMaximizeExpression(A))
	println(mve.DynamicProgrammingMaximizeExpression(A))
	println()

	println("\033[1m\033[31m", "Size 200", "\033[0m")
	A = populateAndShuffleArray(200)
	println(mve.BruteForceMaximizeExpression(A))
	println(mve.DynamicProgrammingMaximizeExpression(A))
	println()

	println("\033[1m\033[31m", "Size 400", "\033[0m")
	A = populateAndShuffleArray(400)
	println(mve.BruteForceMaximizeExpression(A))
	println(mve.DynamicProgrammingMaximizeExpression(A))
	println()

	println("\033[1m\033[31m", "Size 800", "\033[0m")
	A = populateAndShuffleArray(800)
	println(mve.BruteForceMaximizeExpression(A))
	println(mve.DynamicProgrammingMaximizeExpression(A))
	println()

	println("\033[1m\033[31m", "Size 1600", "\033[0m")
	A = populateAndShuffleArray(1600)
	println(mve.BruteForceMaximizeExpression(A))
	println(mve.DynamicProgrammingMaximizeExpression(A))
	println()

}

func populateAndShuffleArray(n int) []int {
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = i
	}

	// shuffle the array
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		A[i], A[j] = A[j], A[i]
	}
	return A
}
