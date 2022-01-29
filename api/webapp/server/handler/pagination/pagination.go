package pagination

// Page ...
type Page struct {
	PageNumber int
	PageSize   int
	Offset     int
}

//CalculateOffset calculation corresponding page number
func (p *Page) CalculateOffset() int {
	p.Offset = (p.PageNumber * p.PageSize) - p.PageSize
	return p.Offset
}
