package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getMiddleKDigits(xn int, k int) int {
	xnStr := strconv.Itoa(xn)
	aux := true
	for (len(xnStr)% 2 != k % 2 || len(xnStr) < k) {
		if aux {
			xnStr = "0"+ xnStr
			aux = false
		} else {
			xnStr = xnStr + "0"
			aux = true
		}
	}
	middle := len(xnStr) / 2
	step := k / 2
	upperBound := middle - step
	lowerBound := middle + step
	if k%2 != 0 {
		lowerBound += 1 
	}
	response, _ := strconv.Atoi(xnStr[upperBound:lowerBound])
	return response
}

func main() {
	var seed int
	flag.IntVar(&seed, "seed", 0, "the seed specifies the start point when a computer generates a random number sequence")
	var k int
	flag.IntVar(&k ,"k", 0, "the number of central digits that will be taken from the previous number")
	var n int
	flag.IntVar(&n, "n", 1, "the number of numbers to generate or the number of iterations")
	
	flag.Parse()
	
	if seed < 1 || k < 1 || n < 1 {
		fmt.Println("missing or invalid arguments")
		os.Exit(1)
	}

	var m = int(math.Pow10(k))
	if seed >= m {
		fmt.Println("the seed must be in the order of 10 ^ k")
		os.Exit(2)
	}

	xn := getMiddleKDigits(int(math.Pow(float64(seed), 2)), k)
	r := float64(xn) / float64(m)
	results := []float64 {r}

	for i := 1; i < n; i++ {
		xn = getMiddleKDigits(int(math.Pow(float64(xn), 2)), k)
		r = float64(xn) / float64(m)
		results = append(results, r)
	}

	fmt.Printf("index\t|\tvalue\n")
	for index, value := range results {
		fmt.Printf("%d\t|\t%v\n", index+1, value)
	}
}