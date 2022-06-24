package square

type Square struct {
	Side float32

	Area float32

	Perimeter float32
}

func (s *Square) AreaOf() float32 {
	s.Area = s.Side * s.Side
	return s.Area
}

func (s *Square) PerimeterOf() float32 {
	s.Perimeter = 4 * s.Side
	return s.Perimeter
}
