package main

import "fmt"

//const N = 3
const N = 2
//const M = 7
const M = 7
//var offset = [N]int {3,0,3}
var offset = [N]int {3,3}
//var skip = [N]int {4,2,1}
var skip = [N]int {4,1}
var permutationtbl [][]int

func main() {
  //fmt.Printf("Hello World\n")
  InitPermutationTable()
	fmt.Println()
	fmt.Println("Populate:")
  Populate()
}

func InitPermutationTable() {
	permutationtbl = make([][]int, N, M)
	for i := 0; i<N; i++ {
	  permutationtbl[i] = make([]int, M)
	  for j := 0; j<M; j++ {
			//permutation := offset[i] + j * skip[i] % M
			permutationtbl[i][j] = (offset[i] + j * skip[i]) % M
    }
  }
	fmt.Println("Permutation Table:")
	for j := 0; j<M; j++ {
	  for i := 0; i<N; i++ {
      fmt.Printf(" %d ", permutationtbl[i][j])
    }
    fmt.Println()
  }
}

func Populate() {
	next  := make([]int, N)
	entry := make([]int, M)
	for i := 0; i<N; i++ {next[i] = 0}
	for j := 0; j<M; j++ {entry[j] = -1}
  var n int = 0
	var c int = 0
	for {
		for i := 0; i<N; i++ {
			fmt.Printf("Case i == %d\n", i)
			c = permutationtbl[i][next[i]]
			fmt.Printf("  c <- %d\n", c)
			for entry[c] >= 0 {
			  fmt.Printf("  entry[%d] == %d\n", c, entry[c])
				next[i] = next[i] + 1
			  fmt.Printf("    next[%d] <- %d\n", i, next[i])
				c = permutationtbl[i][next[i]]
			  fmt.Printf("    c <- %d\n", c)
			}
			entry[c] = i
			fmt.Printf("  entry[%d] <- %d\n", c, entry[c])
			next[i] = next[i] + 1
			fmt.Printf("  next[%d] <- %d\n", i, next[i])
			n = n + 1
			fmt.Printf("  n <- %d\n", n)
		  if n == M {
				break
			}
		}
		//fmt.Printf("n : M = %d %d\n", n, M)
		if n == M {
			break
		}
	}
	for j := 0; j<M; j++ {
		fmt.Printf("%d : %d\n", j, entry[j])
	}
}
