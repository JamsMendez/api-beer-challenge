package pagination

import "github.com/gofiber/fiber/v2"

const (
	DefaultPage    = 1
	DefaultPerPage = 10
	MaxPerPage     = 1000

	KeyPage    = "page"
	KeyPerPage = "per_page"
)

type Pagination struct {
	Items    any     `json:"items"`
	NextPage *uint32 `json:"next_page"`
	Page     uint32  `json:"page"`
	Pages    uint32  `json:"pages"`
	Size     uint32  `json:"size"`
	Total    uint32  `json:"total"`
}

type PaginationParams struct {
	Page    uint32
	PerPage uint32
	Sort    string
	OrderBy string
}

func NewParamsFromFiber(c *fiber.Ctx) *PaginationParams {
	page := c.QueryInt(KeyPage, DefaultPage)
	perPage := c.QueryInt(KeyPerPage, DefaultPerPage)

	if perPage > MaxPerPage {
		perPage = MaxPerPage
	}

	return &PaginationParams{
		Page:    uint32(page),
		PerPage: uint32(perPage),
	}
}

func New(page, size, total uint32) *Pagination {
	var nextPage *uint32

	pages := total / size
	mod := total % size
	if mod > 0 {
		pages++
	}

	if page < pages {
		next := page + 1
		nextPage = &next
	}

	return &Pagination{
		NextPage: nextPage,
		Page:     page,
		Pages:    pages,
		Total:    total,
	}
}
