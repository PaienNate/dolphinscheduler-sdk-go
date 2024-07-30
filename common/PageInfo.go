package common

// PageInfo Dolphin Scheduler分页模型，参考自org.apache.dolphinscheduler.api.utils.PageInfo
// PageInfo Dolphin Scheduler page model, copied from org.apache.dolphinscheduler.api.utils.PageInfo
type PageInfo[T any] struct {
	TotalList   []T // 数据列表 totalList
	Total       int // 总数 total
	TotalPage   int // 总页数 total Page
	PageSize    int // 每页大小 page size
	CurrentPage int // 当前页 current page
	PageNo      int // 页码 pageNo
}

// NewPageInfo 创建一个新的PageInfo实例
// NewPageInfo creates a new PageInfo instance
func NewPageInfo(currentPage, pageSize *int) *PageInfo[any] {
	if currentPage == nil {
		defaultCurrentPage := 1
		currentPage = &defaultCurrentPage
	}
	if pageSize == nil {
		defaultPageSize := 20
		pageSize = &defaultPageSize
	}
	pageNo := (*currentPage - 1) * *pageSize
	return &PageInfo[any]{
		PageSize:    *pageSize,
		CurrentPage: *currentPage,
		PageNo:      pageNo,
	}
}

// GetStart 获取起始位置
// GetStart gets the start position
func (p *PageInfo[T]) GetStart() int {
	return p.PageNo
}

// SetStart 设置起始位置
// SetStart sets the start position
func (p *PageInfo[T]) SetStart(start int) {
	p.PageNo = start
}

// GetTotalList 获取数据列表
// GetTotalList gets the data list
func (p *PageInfo[T]) GetTotalList() []T {
	return p.TotalList
}

// SetTotalList 设置数据列表
// SetTotalList sets the data list
func (p *PageInfo[T]) SetTotalList(totalList []T) {
	p.TotalList = totalList
}

// GetTotal 获取总数
// GetTotal gets the total count
func (p *PageInfo[T]) GetTotal() int {
	return p.Total
}

// SetTotal 设置总数
// SetTotal sets the total count
func (p *PageInfo[T]) SetTotal(total int) {
	p.Total = total
}

// GetTotalPage 获取总页数
// GetTotalPage gets the total page count
func (p *PageInfo[T]) GetTotalPage() int {
	if p.PageSize == 0 {
		p.PageSize = 7
	}
	if p.Total%p.PageSize == 0 {
		if p.Total/p.PageSize == 0 {
			p.TotalPage = 1
		} else {
			p.TotalPage = p.Total / p.PageSize
		}
	} else {
		p.TotalPage = p.Total/p.PageSize + 1
	}
	return p.TotalPage
}

// SetTotalPage 设置总页数
// SetTotalPage sets the total page count
func (p *PageInfo[T]) SetTotalPage(totalPage int) {
	p.TotalPage = totalPage
}

// GetPageSize 获取每页大小
// GetPageSize gets the page size
func (p *PageInfo[T]) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 7
	}
	return p.PageSize
}

// SetPageSize 设置每页大小
// SetPageSize sets the page size
func (p *PageInfo[T]) SetPageSize(pageSize int) {
	p.PageSize = pageSize
}

// GetCurrentPage 获取当前页
// GetCurrentPage gets the current page
func (p *PageInfo[T]) GetCurrentPage() int {
	if p.CurrentPage <= 0 {
		p.CurrentPage = 1
	}
	return p.CurrentPage
}

// SetCurrentPage 设置当前页
// SetCurrentPage sets the current page
func (p *PageInfo[T]) SetCurrentPage(currentPage int) {
	p.CurrentPage = currentPage
}
