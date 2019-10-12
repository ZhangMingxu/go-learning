package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	for i := 0; i < dy; i++ {
		var s []uint8
		for j := 0; j < dx; j++ {
			s = append(s, uint8(i*j))
		}
		result = append(result, s)
	}
	return result
}

func main() {
	pic.Show(Pic)
}
