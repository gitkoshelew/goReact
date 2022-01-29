package pagination

type Page struct{
	PageNumber int
	PageSize int
	Offset int 
}

//offset calculation corresponding page number
func (p *Page) CalculateOffset() int{
	p.Offset = (p.PageNumber * p.PageSize) - p.PageSize
	return p.Offset
}
