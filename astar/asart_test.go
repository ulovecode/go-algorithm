package astar

import (
	"fmt"
	"testing"
	"time"
)

func TestAStar_FindPath(t *testing.T) {
	t1 := time.Now() // get current time
	start := NewNode(0,0)
	end := NewNode(4,5)
	searchMap := [][]int{
		{0,0,0,0,0,0,0,0,0},//0
		{0,0,0,0,0,-1,0,-1,0},//1
		{0,0,0,0,0,0,0,0,0},//2
		{0,0,0,-1,-1,-1,-1,-1,0},//3
		{0,0,0,-1,0,999,0,0,0},//4
		{0,0,0,0,-1,0,0,0,0},//5
		{0,0,0,0,0,0,0,0,0},//6
	}
	star := NewAStar(start, end, searchMap)
	star.FindPath()
	star.PrintNode()
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}