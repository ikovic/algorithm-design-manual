package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x, y int
}

type Points []Point

func (points Points) print(name string) {
	fmt.Printf("Result for %s: \n", name)
	for i := range points {
		fmt.Printf("{%d, %d}\n", points[i].x, points[i].y)
	}
}

func (source Point) getDistance(destination Point) float64 {
	first := math.Pow(float64(destination.x-source.x), 2)
	second := math.Pow(float64(destination.y-source.y), 2)
	return math.Sqrt(first + second)
}

/*
	Bad heuristic: NearestNeighbor(P )

	Pick and visit an initial point p0 from P
	p = p0
	i = 0
	While there are still unvisited points
		i=i+1
		Select pi to be the closest unvisited point to pi−1
		Visit pi
	Return to p0 from pn−1
*/
func nearestNeighbour(input []Point) []Point {
	s := input[:]
	p := input[0]
	visited := []Point{p}
	i := 0

	for len(visited) != len(input) {
		i++

		sort.Slice(s, func(i, j int) bool {
			return p.getDistance(s[i]) < p.getDistance(s[j])
		})

		s = s[1:]
		p = s[0]

		visited = append(visited, p)
	}

	return append(visited, visited[0])
}

/*
	Bad heuristic: ClosestPair(P)
		Let n be the number of points in set P.
		For i = 1 to n − 1 do
			d=∞
			For each pair of endpoints (s, t) from distinct vertex chains
				if dist(s,t)≤d then sm=s, tm=t, and d=dist(s,t)
			Connect (sm, tm) by an edge
		Connect the two endpoints by an edge
*/
func closestPair(input []Point) []Point {
	return input
}

/*
	Problem: Robot Tour Optimization
	Input: A set S of n points in the plane.
	Output: What is the shortest cycle tour that visits each point in the set S?
*/
func main() {
	input := Points{{0, 0}, {3, 0}, {2, 0}, {5, 0}}

	var nearestNeighbourSolution, closestPairSolution Points
	nearestNeighbourSolution = nearestNeighbour(input)
	closestPairSolution = closestPair(input)

	nearestNeighbourSolution.print("nearest neighbour")
	closestPairSolution.print("closest pair")
}
