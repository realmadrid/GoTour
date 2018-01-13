package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
	}
	for i, j := range p {
		for k := range j {
			j[k] = (uint8)(i+k) / 2
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
