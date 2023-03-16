# go-matmul-parallel

## 1. How to Run
- go build
- ./go-matmul-parallel
    - Input a matrix size 
    - ex) Matrix Size: 100

## 2. Implementation Detail
- Go의 concurrency 지원을 위해 개발된 goroutine 테스트 프로그램
- Matrix multiplication 연산을 goroutine을 통해서 병렬화(parallelization)
- func matmul_seq(A []int, B []int, C []int, N int)
    - Matrix multiplication의 순차 실행 버전
- func matmul_partial(A []int, B []int, C []int, N int, idx int, num_thread int)
    - Matrix multiplication의 병렬 실행 버전
    - goroutine을 통해 병렬 실행함
        - M=runtime.NumCPU()개의 goroutine 실행
        - runtime.GOMAXPROCS(M)을 이용해 M개의 process 실행
        - 이를 통해 Concurrency가 아닌 Parallelism 방식으로 구현함

## 3. Sample Result
```
$ ./go-matmul-parallel
Matrix Size: 1000
Sequential Matrix Multiplication started...
 - ended.
 - elapsed time: 1.328565318s
Parallel Matrix Multiplication started...
 - ended.
 - elapsed time: 202.288343ms
Checking the results succeeded!!
```