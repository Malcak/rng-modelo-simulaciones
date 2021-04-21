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
	var seedOne int
	flag.IntVar(&seedOne, "seedOne", 0, "the seed specifies the start point when a computer generates a random number sequence")
	var seedTwo int
	flag.IntVar(&seedTwo, "seedTwo", 0, "the seed specifies the start point when a computer generates a random number sequence")
	var k int
	flag.IntVar(&k ,"k", 0, "the number of central digits that will be taken from the previous number")
	var n int
	flag.IntVar(&n, "n", 1, "the number of numbers to generate or the number of iterations")
	
	flag.Parse()
	
	if seedOne < 1 || seedTwo < 1 || k < 1 || n < 1 {
		fmt.Println("missing or invalid arguments")
		os.Exit(1)
	}

	var m = int(math.Pow10(k))
	if seedOne >= m || seedTwo >= m {
		fmt.Println("the seed must be in the order of 10 ^ k")
		os.Exit(2)
	}

	xn0 := getMiddleKDigits(seedOne * seedTwo, k)
	r := float64(xn0) / float64(m)
	results := []float64 {r}

	xn1 := getMiddleKDigits(seedTwo * xn0, k)
	r = float64(xn1) / float64(m)
	results = append(results, r)

	var temp int

	for i := 2; i < n; i++ {

		temp = xn1
		
		xn1 = getMiddleKDigits(xn0 * xn1, k)
		r = float64(xn1) / float64(m)
		results = append(results, r)

		xn0 = temp
	}

	fmt.Printf("index\t|\tvalue\n")
	for index, value := range results {
		fmt.Printf("%d\t|\t%v\n", index+1, value)
	}
}