package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/* ******************************************************************************
 * Seqential Matrix Multiplication C = A * B
 * ******************************************************************************/
func matmul_seq(A []int, B []int, C []int, N int) {

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			sum := 0
			for k := 0; k < N; k++ {
				sum += A[y*N+k] * B[k*N+x]
			}
			C[y*N+x] = sum
		}
	}
}

/* ******************************************************************************
 * Seqential Matrix Multiplication C = A * B
 * ******************************************************************************/
func matmul_partial(A []int, B []int, C []int, N int, idx int, num_thread int) {

	y_start := (N / num_thread) * idx
	y_end := (N / num_thread) * (idx + 1)
	if idx+1 == num_thread {
		y_end = N
	}

	for y := y_start; y < y_end; y++ {
		for x := 0; x < N; x++ {
			sum := 0
			for k := 0; k < N; k++ {
				sum += A[y*N+k] * B[k*N+x]
			}
			C[y*N+x] = sum
		}
	}
}

func main() {
	/* ******************************************************************************
	 * Configurate Matrices
	 * ******************************************************************************/
	var N int
	fmt.Printf("Matrix Size: ")
	fmt.Scanln(&N)

	/* ******************************************************************************
	 * Make arrays A, B, C
	 * ******************************************************************************/
	A := make([]int, N*N)
	B := make([]int, N*N)
	C_seq := make([]int, N*N)
	C_par := make([]int, N*N)

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			A[y*N+x] = rand.Int()%11 - 5
			B[y*N+x] = rand.Int()%11 - 5
		}
	}

	/* ******************************************************************************
	 * Run matmul squentially C = A * B
	 * ******************************************************************************/
	fmt.Printf("Sequential Matrix Multiplication started...\n")
	start_seq := time.Now()
	matmul_seq(A, B, C_seq, N)
	fmt.Printf(" - ended.\n")
	fmt.Printf(" - elapsed time: %v\n", time.Since(start_seq))

	/* ******************************************************************************
	 * Run matmul parallelly via all CPUs C = A * B
	 * ******************************************************************************/
	fmt.Printf("Parallel Matrix Multiplication started...\n")
	num_thread := runtime.NumCPU()
	if N < num_thread {
		num_thread = N
	}
	runtime.GOMAXPROCS(num_thread)

	start_par := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < num_thread; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			matmul_partial(A, B, C_par, N, i, num_thread)
		}(i)
	}
	wg.Wait()
	fmt.Printf(" - ended.\n")
	fmt.Printf(" - elapsed time: %v\n", time.Since(start_par))

	/* ******************************************************************************
	 * Check results
	 * ******************************************************************************/
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if C_seq[y*N+x] != C_par[y*N+x] {
				log.Fatalf("[ERROR] Checking the results failed at [%v,%v] (%v != %v)", y, x, C_seq[y*N+x], C_par[y*N+x])
			}
		}
	}
	fmt.Print("Checking the results succeeded!!\n")
}

/* ******************************************************************************
 * Helper function
 * ******************************************************************************/
func print_matrix(C []int, N int) {
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			fmt.Printf("%v ", C[y*N+x])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

}
