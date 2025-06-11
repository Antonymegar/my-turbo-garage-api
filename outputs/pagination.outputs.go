package outputs

// Pagination ...
type Pagination struct {
	Page     int64 `form:"page,default=1"`
	PageSize int64 `form:"pageSize,default=20" binding:"max=100"`
}

// getPage ...
func (p *Pagination) getPage() int {
	page := int(p.Page)
	if page == 0 {
		return 1
	}
	return page
}

// GetLimit ...
func (p *Pagination) GetLimit() int {
	size := int(p.PageSize)
	if size == 0 {
		return 20
	}
	return size
}

// GetOffset ...
func (p *Pagination) GetOffset() int {
	return (p.getPage() - 1) * p.GetLimit()
}
