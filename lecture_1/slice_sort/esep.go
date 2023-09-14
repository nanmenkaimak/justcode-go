package main

type sliceSorting struct {
	values []int
}

func (s *sliceSorting) sorting() []int {
	for i := 0; i < len(s.values)-1; i++ {
		for j := 0; j < len(s.values)-i-1; j++ {
			if s.values[j] > s.values[j+1] {
				s.values[j], s.values[j+1] = s.values[j+1], s.values[j]
			}
		}
	}
	return s.values
}
