package ktc

type Size struct {
	MinWidth  int
	MinHeight int

	Width  int
	Height int
}

func (s *Size) GetWidth() int {
	// Default
	if s.Width == 0 {
		return 1024
	}
	// Min align
	if s.Width < s.MinWidth {
		return s.MinWidth
	}
	// Return current value
	return s.Width
}

func (s *Size) GetHeight() int {
	// Default
	if s.Height == 0 {
		return 512
	}
	// Min align
	if s.Height < s.MinHeight {
		return s.MinHeight
	}
	// Return current value
	return s.Height
}

func (r *Size) Resize(width, height int) {
	r.Width = width
	r.Height = height
	// Min values align
	if r.Width < r.MinWidth {
		r.Width = r.MinWidth
	}
	if r.Height < r.MinHeight {
		r.Height = r.MinHeight
	}
}
