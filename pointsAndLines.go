package main

import "fmt"
import "strconv"


type Point struct {
	X int
	Y int
}

type Line struct {
	K float64
	L float64
}

var points =  map[Point]Point{
	{0,1}: {1,1},
	{0,0}: {1,1},
	{0,2}: {1,1},
	{0,4}: {1,1},
	{0,-1}: {1,1},
}
func main() {

	for a,b := range points {

		fmt.Printf("First point:")
		printPoint(a)
		fmt.Printf("Second point:")
		printPoint(b)

		abLine := getLine(a,b)
		printLine(abLine)
		fmt.Printf("___\n\n")
	}
}

func getLine(a Point, b Point) Line {

	var line Line
	line.K = float64(a.Y-b.Y)/float64(a.X-b.X)
	line.L = float64(a.Y) - (line.K * float64(a.X))
	return line
}

func printPoint(point Point) {
	fmt.Printf("(%v,%v)\n", point.X, point.Y)
}

func printLine(line Line) {

	var result string;
	result = "Line: y="

	switch { 
		case line.K == 0:
			result = result + strconv.FormatFloat(line.L, 'f', -1, 64)
		case line.K != 1:
			result = result + strconv.FormatFloat(line.K, 'f', -1, 64)
			fallthrough
		default:
			result = result + "x"
			switch {
				case line.L > 0:
					result = result + "+" + strconv.FormatFloat(line.L, 'f', -1, 64)
				case line.L < 0:
					result = result + strconv.FormatFloat(line.L, 'f', -1, 64)
			}
	}
	fmt.Println(result)
}
