package models

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit" form:"limit"`
	Page       int         `json:"page" form:"page"`
	TotalRows  int         `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = int(totalRows)
	totalPages := int((pagination.TotalRows / pagination.GetLimit()))
	pagination.TotalPages = totalPages
	if pagination.TotalRows == 0 {
		pagination.TotalPages = 0
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
