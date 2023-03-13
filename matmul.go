package main

import (
	"fmt"
	"math/rand"
)

func matmul_seq(A []int, B []int, N int) []int {

	C := make([]int, N*N)

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			sum := 0
			for k := 0; k < N; k++ {
				sum += A[y*N+k] * B[k*N+x]
			}
			C[y*N+x] = sum
		}
	}

	return C
}

func main() {

	// Size of Matrices:(M*M)
	var N int
	fmt.Printf("Matrix Size: ")
	fmt.Scanln(&N)

	// Make arrays A, B, C
	A := make([]int, N*N)
	B := make([]int, N*N)

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			A[y*N+x] = rand.Int()%11 - 5
			B[y*N+x] = rand.Int()%11 - 5
		}
	}

	C := matmul_seq(A, B, N)
	print_matrix(A, N)
	print_matrix(B, N)
	print_matrix(C, N)
}

// Helper function
func print_matrix(C []int, N int) {
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			fmt.Printf("%d ", C[y*N+x])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}
