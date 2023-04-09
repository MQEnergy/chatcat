package cpaginator

import (
	"chatcat/backend/config"
	"chatcat/backend/pkg/chelp"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"math"
)

type PageBuilder struct {
	DB       *gorm.DB
	Cfg      *config.Conf
	Model    interface{} // model struct
	Preloads []string    // 预加载
	Fields   []string    // 查询字段
}

type Page struct {
	List        interface{} `json:"list"`         // 查询的列表
	CurrentPage int         `json:"current_page"` // 当前页
	Total       int64       `json:"total"`        // 查询记录总数
	LastPage    int         `json:"last_page"`    // 最后一页
	PerPage     int         `json:"per_page"`     // 每页条数
}

func NewBuilder(db *gorm.DB, cfg *config.Conf) *PageBuilder {
	return &PageBuilder{
		DB:  db,
		Cfg: cfg,
	}
}

// WithField 查询单表的字段 和 过滤字段 不能与WithFields方法同用
func (pb *PageBuilder) WithField(fields []string) *PageBuilder {
	fieldList := filterFields(pb.Model, fields)
	pb.Fields = fieldList
	pb.DB.Select(pb.Fields)
	return pb
}

// filterFields 过滤查询字段
func filterFields(model interface{}, fields []string) []string {
	var fieldList []string
	if len(fields) == 1 {
		if fields[0] != "_omit" && fields[0] != "_select" {
			fieldList = fields
		}
	} else {
		switch fields[len(fields)-1] {
		case "_omit":
			fields = fields[:len(fields)-1]
			_fields, _ := chelp.GetStructColumnName(model, 1)
			fieldList, _ = lo.Difference[string](_fields, fields)
		case "_select":
			fieldList = fields[:len(fields)-1]
		default:
			fieldList = fields[:]
		}
	}
	return fieldList
}

// WithModel 查询的model struct
func (pb *PageBuilder) WithModel(model interface{}) *PageBuilder {
	pb.Model = model
	pb.DB = pb.DB.Model(&model)
	return pb
}

// WithOrderBy 排序
func (pb *PageBuilder) WithOrderBy(orderBy interface{}) *PageBuilder {
	pb.DB = pb.DB.Order(orderBy)
	return pb
}

// WithPreloads 多表关联查询主动预加载（无条件）
func (pb *PageBuilder) WithPreloads(querys []string) *PageBuilder {
	pb.Preloads = querys
	return pb
}

// WithPreload 关联查询主动预加载（可传条件）
func (pb *PageBuilder) WithPreload(query string, args ...interface{}) *PageBuilder {
	pb.DB.Preload(query, args...)
	return pb
}

// WithCondition 查询条件
func (pb *PageBuilder) WithCondition(query interface{}, args ...interface{}) *PageBuilder {
	pb.DB.Where(query, args...)
	return pb
}

// Pagination 分页查询
func (pb *PageBuilder) Pagination(dst interface{}, currentPage, pageSize int) (Page, error) {
	query := pb.DB
	page := pb.ParsePage(currentPage, pageSize)
	offset := (page.CurrentPage - 1) * page.PerPage

	// 查询总数
	if err := query.Count(&page.Total).Error; err != nil {
		return page, err
	}
	// 预加载
	if len(pb.Preloads) > 0 {
		for _, preload := range pb.Preloads {
			query.Preload(preload)
		}
	}
	// 计算总页数
	if page.Total > int64(page.PerPage) {
		page.LastPage = int(math.Ceil(float64(page.Total) / float64(page.PerPage)))
	}
	// 判断总数跟最后一页的关系
	if page.CurrentPage <= page.LastPage {
		if err := query.Limit(page.PerPage).Offset(offset).Find(dst).Error; err != nil {
			return page, err
		}
	}
	page.List = dst
	return page, nil
}

// ParsePage 分页超限设置和格式化
func (pb *PageBuilder) ParsePage(currentPage, pageSize int) Page {
	var page Page
	// 返回每页数量
	page.PerPage = pageSize
	// 返回当前页码
	page.CurrentPage = currentPage

	if currentPage < 1 {
		page.CurrentPage = 1
	}
	if pageSize < 1 {
		page.PerPage = pb.Cfg.App.DefaultPageSize
	}
	if pageSize > pb.Cfg.App.MaxPageSize {
		page.PerPage = pb.Cfg.App.MaxPageSize
	}
	if page.LastPage < 1 {
		page.LastPage = 1
	}
	return page
}
