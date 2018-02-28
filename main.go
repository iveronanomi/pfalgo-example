package main

import (
	"fmt"

	"github.com/iveronanomi/pfalgo"
)

func main() {
	sg, start, end := getGrid()
	//cf := pfalgo.BreadthFirstSearch(sg, start, end, callback)
	//cf := pfalgo.GreedyBreadthFirstSearch(sg, start, end, callback)
	//cf, _ := pfalgo.DijkstraSearch(sg, start, end, callback)
	cf, _ := pfalgo.AStarSearch(sg, start, end, callback)
	fmt.Println(pfalgo.String(sg))
	rp := pfalgo.ReconstructPath(cf, start, end, false)
	drawReconstructedPath(rp)
}

func callback (g *pfalgo.GridGraph, start, current pfalgo.INode) {
	//mark graph node as visited
	if !current.Equal(start) {
		g.Visit(current.Position())
	}
}

func getGrid() (*pfalgo.GridGraph, pfalgo.Node, pfalgo.Node) {
	//Create Grid
	w, h := 30, 15
	sg := pfalgo.NewSquareGrid(
		uint32(w),
		uint32(h),
		pfalgo.LinearWalk,
		pfalgo.NewGifGraph(w, h, "out/out.gif"),
	)
	startX, startY := 24, 12
	endX, endY := 0, 10

	//Add walls to grid
	sg.AddWall(3, 3, 9, 2)
	sg.AddWall(13, 4, 11, 2)
	sg.AddWall(21, 0, 7, 2)
	sg.AddWall(21, 5, 2, 5)
	sg.Start(startX, startY) //only for a draw
	sg.Target(endX, endY) //only for a draw

	return sg, pfalgo.Node{X:startX,Y:startY}, pfalgo.Node{X:endX,Y:endY}
}

func drawReconstructedPath(rp []pfalgo.INode) {
	sg, start, end := getGrid()
	for _, v := range rp {
		if v.Equal(start) || v.Equal(end) {
			continue
		}
		sg.Visit(v.Position())
	}
	fmt.Println(pfalgo.String(sg))
}
