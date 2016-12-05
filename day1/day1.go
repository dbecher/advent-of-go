package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) distance() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func (p Point) add(x, y int) Point {
	return Point{x: p.x + x, y: p.y + y}
}

func (p Point) lineTo(other Point) Edge {
	return Edge{start: p, end: other}
}

type Edge struct {
	start, end Point
}

func (e Edge) intersects(other Edge) (bool, Point) {
	eStart, eEnd, otherStart, otherEnd := e.start, e.end, other.start, other.end
	// for the purposes of this we dont care about colinear lines (since there
	// will always be a perpendicular one that intersects), so eliminate any that
	// are parallel
	if (eStart.x == eEnd.x && otherStart.x == otherEnd.x) ||
		(eStart.y == eEnd.y && otherStart.y == otherEnd.y) {
		return false, Point{}
	}
	// also for the purposes of this we don't care about the direction, so make the
	// lines all go the same way
	if eStart.x > eEnd.x || eStart.y > eEnd.y {
		eStart, eEnd = eEnd, eStart
	}
	if otherStart.x > otherEnd.x || otherStart.y > otherEnd.y {
		otherStart, otherEnd = otherEnd, otherStart
	}
	// e should be the vertical line, other the horizontal
	if eStart.x != eEnd.x {
		eStart, eEnd, otherStart, otherEnd = otherStart, otherEnd, eStart, eEnd
	}
	if (eStart.x < otherStart.x || eEnd.x > otherEnd.x) || (eStart.y > otherStart.y || eEnd.y < otherStart.y) {
		return false, Point{}
	}
	return true, Point{eStart.x, otherStart.y}
}

type Move struct {
	// +1: right
	// -1: left
	direction int
	distance  int
}

func main() {
	moves := getInput("input.txt")

	allEdges := make([]Edge, len(moves))
	currentPosition := Point{}
	isX := true
	xDirection, yDirection := 1, 1

	for i, move := range moves {
		var to Point
		if isX {
			direction := move.direction * xDirection
			to = currentPosition.add(move.distance*direction, 0)
			yDirection = -1 * direction
		} else {
			direction := move.direction * yDirection
			to = currentPosition.add(0, move.distance*direction)
			xDirection = direction
		}

		// check if we've been here
		newEdge := currentPosition.lineTo(to)
		var intersectsAt Point
		foundIntersection := false
		// skip the most recently added edge (that is, i - 1), since we know it will
		// always intersect this one and is not the answer
		for j := 0; j < i-1; j++ {
			foundIntersection, intersectsAt = newEdge.intersects(allEdges[j])
			if foundIntersection {
				break
			}
		}
		if foundIntersection {
			currentPosition = intersectsAt
			break
		}

		isX = !isX
		currentPosition = to
		allEdges[i] = newEdge
	}

	// result is currentPosition
	fmt.Println(currentPosition.distance())
}

func getInput(filename string) []Move {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.TrimSpace(string(dat)), ", ")
	moves := make([]Move, len(input))
	for i, m := range input {
		tuple := strings.SplitN(m, "", 2)
		direction := 1
		if tuple[0] == "L" {
			direction = -1
		}
		distance, err := strconv.Atoi(tuple[1])
		if err != nil {
			panic(err)
		}
		moves[i] = Move{direction: direction, distance: distance}
	}
	return moves
}
