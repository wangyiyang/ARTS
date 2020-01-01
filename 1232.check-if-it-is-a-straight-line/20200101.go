package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	x1 := coordinates[1][0] - coordinates[0][0]
	y1 := coordinates[1][1] - coordinates[0][1]
	for i := 2; i <= len(coordinates); i++ {
		x2 := coordinates[i][0] - coordinates[0][0]
		y2 := coordinates[i][1] - coordinates[0][1]
		if x1*y2 != x2*y1 {
			return false
		}
	}
	return true
}

func main() {

}
