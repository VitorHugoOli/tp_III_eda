package mve

import "math"

// Brute force for Maximize the value of an expression, given an array A, maximize value of expression (A[s] - A[r] + A[q] - A[p]), where p, q, r, and s are indices of the array and s > r > q > p
func BruteForceMaximizeExpression(A []int) int {
	comparison := 0
	max := -1000000000
	for p := 0; p < len(A); p++ {
		for q := p + 1; q < len(A); q++ {
			for r := q + 1; r < len(A); r++ {
				for s := r + 1; s < len(A); s++ {
					comparison++
					if A[s]-A[r]+A[q]-A[p] > max {
						max = A[s] - A[r] + A[q] - A[p]
					}
				}
			}
		}
	}
	println("\033[1m\033[34mBrute: ", comparison, "\033[0m")
	return max
}

// Dynamic programming for Maximize the value of an expression, given an array A, maximize value of expression (A[s] - A[r] + A[q] - A[p]), where p, q, r, and s are indices of the array and s > r > q > p
func DynamicProgrammingMaximizeExpression(A []int) int {
	comparison := 0
	n := len(A)
	// check if array is bigger than 4 length
	if len(A) < 4 {
		panic("Array length must be bigger than 4")
	}
	comparison++

	// initialize 4 lookup tables for first, second, third, fourth and populate it with the lowest possible value
	first := make([]int, n+1)
	second := make([]int, n)
	third := make([]int, n-1)
	fourth := make([]int, n-2)
	for i := 0; i < n-3; i++ {
		comparison++
		first[i] = -1000000000
		second[i] = -1000000000
		third[i] = -1000000000
		fourth[i] = -1000000000
	}
	first[n-2] = -1000000000
	second[n-2] = first[n-2]
	third[n-2] = second[n-2]

	first[n-1] = -1000000000
	second[n-1] = first[n-1]

	first[n] = second[n-1]

	// `first[]` stores the maximum value of `A[l]`
	for i := n - 1; i >= 0; i-- {
		comparison++
		first[i] = int(math.Max(float64(first[i+1]), float64(A[i])))
	}

	// `second[]` stores the maximum value of `A[l] - A[k]`
	for i := n - 2; i >= 0; i-- {
		comparison++
		second[i] = int(math.Max(float64(second[i+1]), float64(first[i+1]-A[i])))
	}

	// `third[]` stores the maximum value of `A[l] - A[k] + A[j]`
	for i := n - 3; i >= 0; i-- {
		comparison++
		third[i] = int(math.Max(float64(third[i+1]), float64(second[i+1]+A[i])))
	}

	// `fourth[]` stores the maximum value of `A[l] - A[k] + A[j] - A[i]`
	for i := n - 4; i >= 0; i-- {
		comparison++
		fourth[i] = int(math.Max(float64(fourth[i+1]), float64(third[i+1]-A[i])))
	}

	//print bold and blue the total number of comparisons
	println("\033[1m\033[34mDynamic: ", comparison, "\033[0m")
	return fourth[0]
}
