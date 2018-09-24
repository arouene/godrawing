package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	imageWidth    = 3000
	imageHeight   = 3000
	length        = 100.0
	border        = 10.0
	hexagoneAngle = (1.0 / 3.0) * math.Pi
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	clearImage(img)

	const width = border + length
	const height = border + length*0.87 // sin(60)
	const xOffset = width + width*0.5   // cos(60)
	const yOffset = height

	// draw two grids of hexagones
	drawGrid(img, width, height, 0, 0, color.Black)
	drawGrid(img, width, height, xOffset, yOffset, color.Gray{128})

	png.Encode(os.Stdout, img)
}

func drawGrid(img *image.RGBA, width, height, xOffset, yOffset int, c color.Color) {
	for i := -height + yOffset; i < imageHeight+height; i += 2 * height {
		for j := -width + xOffset; j < imageWidth+width; j += 3 * width {
			drawHexagone(img, j, i, length, c)
		}
	}
}

func drawHexagone(img *image.RGBA, x, y int, l float64, c color.Color) {
	nextHexagonePoint := nextPolygonePoint(x, y, l, hexagoneAngle)
	x0, y0, _ := nextHexagonePoint()
	x1, y1, _ := nextHexagonePoint()
	for end := false; !end; x1, y1, end = nextHexagonePoint() {
		drawLine(img, x0, y0, x1, y1, c)
		x0, y0 = x1, y1
	}
}

func nextPolygonePoint(x, y int, l, t float64) func() (int, int, bool) {
	angle := 0.0
	return func() (x0 int, y0 int, end bool) {
		x0 = x + int(l*math.Cos(angle))
		y0 = y + int(l*math.Sin(angle))
		angle += t
		if angle > (2*math.Pi)+t+0.1 {
			end = true
		}
		return
	}
}

// clearImage is an optimisation of drawRectangle where we clear
// the full image with white color, about 3 times faster
func clearImage(img *image.RGBA) {
	for i := 0; i < len(img.Pix); i++ {
		img.Pix[i] = 0xff
	}
}

func drawRectangle(img *image.RGBA, x0, y0, x1, y1 int, c color.Color) {
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	for i := x0; i < x1; i += sx {
		for j := y0; j < y1; j += sy {
			img.Set(i, j, c)
		}
	}
}

// drawLine draw a line with the bresenham algorithm
func drawLine(img *image.RGBA, x0, y0, x1, y1 int, c color.Color) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		img.Set(x0, y0, c)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
