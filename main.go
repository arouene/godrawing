package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const length = 100.0
const border = 10.0

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 3000, 3000))
	drawRectangle(img, 0, 0, 3000, 3000, color.White)

	const width = border + length
	const height = border + length*0.87 // sin(60)
	const xOffset = width + width*0.5
	const yOffset = height

	// draw two grids of hexagones
	drawGrid(img, width, height, 0, 0, color.Black)
	drawGrid(img, width, height, xOffset, yOffset, color.Gray{128})

	png.Encode(os.Stdout, img)
}

func drawGrid(img *image.RGBA, width, height, xOffset, yOffset int, c color.Color) {
	for i := height + yOffset; i < 3000; i += 2 * height {
		for j := width + xOffset; j < 3000; j += 3 * width {
			drawHexagone(img, j, i, length, c)
		}
	}
}

func drawHexagone(img *image.RGBA, x, y int, l float64, c color.Color) {
	const angle = (1.0 / 3.0) * math.Pi
	x0, y0 := nextHexagonePoint(x, y, l, 0)
	x1, y1 := 0, 0
	for i := angle; i <= (2*math.Pi)+0.1; i += angle {
		x1, y1 = nextHexagonePoint(x, y, l, i)
		drawLine(img, x0, y0, x1, y1, c)
		x0, y0 = x1, y1
	}
}

func nextHexagonePoint(x, y int, l, t float64) (x0 int, y0 int) {
	x0 = x + int(l*math.Cos(t))
	y0 = y + int(l*math.Sin(t))
	return
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
