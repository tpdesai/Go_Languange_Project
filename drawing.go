package main

import (
	"canvas"
	//"image"
)


//DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
//object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(b CheckerBoard, canvasWidth int) canvas.Canvas {
	// set a new square canvas
	height := len(b) * canvasWidth
	width := len(b[0]) * canvasWidth
	c := canvas.CreateNewCanvas(width, height)



	// range over all the bodies and draw them.
	for i := range b {
		for j := range b[i] {
			if b[i][j] == 0 {
				c.SetFillColor(canvas.MakeColor(0,0,0))
		 }else if b[i][j] == 1 {
				c.SetFillColor(canvas.MakeColor(85,85,85))
		 }else if b[i][j] == 2 {
			 	c.SetFillColor(canvas.MakeColor(170,170,170))
		 }else if b[i][j] == 3{
			 	c.SetFillColor(canvas.MakeColor(255,255,255))
		 }
	
	x := i * canvasWidth
	y := j * canvasWidth
	c.ClearRect(x, y, x+canvasWidth, y+canvasWidth)
}
}

	// we want to return an image!
	return c
}
