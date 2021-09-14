package sports

type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *shirt) GetLogo() string {
	return s.logo
}

func (s *shirt) SetSize(size int) {
	s.size = size
}

func (s *shirt) GetSize() int {
	return s.size
}