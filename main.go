package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	drawRectangle(img, 0, 0, 300, 300, color.White)
	drawHexagone(img, 150, 150, 100, color.Black)
	png.Encode(os.Stdout, img)
}

func drawHexagone(img *image.RGBA, x, y int, l float64, c color.Color) {
	const angle = (1.0 / 3.0) * math.Pi
	x0, y0 := nextPoint(x, y, l, 0)
	x1, y1 := x0, y0
	for i := angle; i < (2*math.Pi)+1; i += angle {
		x1, y1 = nextPoint(x, y, l, i)
		drawLine(img, x0, y0, x1, y1, c)
		x0, y0 = x1, y1
	}
}

func nextPoint(x, y int, l, t float64) (x0 int, y0 int) {
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
