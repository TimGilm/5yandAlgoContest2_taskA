package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	k, _ := strconv.Atoi(s.Text())
	var arrX []int
	var arrY []int

	for i := 0; i < k; i++ {
		var point [2]int
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		line = strings.TrimSpace(line)
		arr := strings.Split(line, " ")
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])

		arrX = append(arrX, x)
		arrY = append(arrY, y)
		point[0] = x
		point[1] = y
	}
	minX := findMinX(arrX)
	maxX := findMaxX(arrX)
	minY := findMinY(arrY)
	maxY := findMaxY(arrY)

	fmt.Print(minX, maxX, minY, maxY)
}
func findMinX(arrX []int) int {
	minX := arrX[0]
	for i, _ := range arrX {
		if arrX[i] < minX {
			minX = arrX[i]
		}
	}
	return minX
}
func findMaxX(arrX []int) int {
	maxX := arrX[0]
	for i, _ := range arrX {
		if arrX[i] > maxX {
			maxX = arrX[i]
		}
	}
	return maxX
}
func findMinY(arrY []int) int {
	minY := arrY[0]
	for i, _ := range arrY {
		if arrY[i] < minY {
			minY = arrY[i]
		}
	}
	return minY
}
func findMaxY(arrY []int) int {
	maxY := arrY[0]
	for i, _ := range arrY {
		if arrY[i] > maxY {
			maxY = arrY[i]
		}
	}
	return maxY
}
