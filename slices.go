package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
		for x := range s[i] {
			s[i][x] = uint8(i+x)
		}
	}
	fmt.Println(s)
	return s
}

func main() {
	pic.Show(Pic)
}

