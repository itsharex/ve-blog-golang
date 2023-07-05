package request

import (
	"context"
	"fmt"
)

// 请求上下文,一般存放请求头参数
type Context struct {
	context.Context `json:"-" header:"-"`
	Token           string `json:"token" header:"token" example:""`
	UID             int    `json:"uid" header:"-" example:""`
	Username        string `json:"username" header:"-" example:""`
	Ip              string `json:"ip" header:"-" example:""`
	IpSource        string `json:"ip_source" header:"-" example:""`
}

func (s *Context) GetContext() context.Context {
	return s.Context
}

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page       int          `json:"page" form:"page"`             // 页码
	PageSize   int          `json:"page_size" form:"page_size"`   // 每页大小
	Order      string       `json:"order" form:"order"`           // 排序关键词
	OrderKey   string       `json:"order_key" form:"order_key"`   // 排序 asc|desc
	Orders     []*Order     `json:"orders" form:"orders"`         // 排序
	Conditions []*Condition `json:"conditions" form:"conditions"` // 使用条件语句查询
}

type Order struct {
	Field string `json:"field"` // 表字段
	Rule  string `json:"rule"`  // 排序规则 asc|desc
}

// 字段，关键字，匹配规则
type Condition struct {
	Flag  string      `json:"flag"`  // 标识 and、or,默认and
	Field string      `json:"field"` // 表字段
	Rule  string      `json:"rule"`  // 规则 =、like、in
	Value interface{} `json:"value"` // 值
}

func (page *PageInfo) Limit() int {
	if page.PageSize == 0 {
		page.PageSize = 10
	}
	return page.PageSize
}

func (page *PageInfo) Offset() int {

	return (page.Page - 1) * page.Limit()
}

// 排序语句
func (page *PageInfo) OrderClause() string {
	if len(page.Orders) == 0 {
		return ""
	}

	var query string
	var flag string
	for i, order := range page.Orders {
		if i == 0 {
			flag = ""
		} else {
			flag = ","
		}
		query += fmt.Sprintf("%s %s %s", flag, order.Field, order.Rule)
	}

	return query
}

// 条件语句
func (page *PageInfo) WhereClause() (string, []interface{}) {
	if len(page.Conditions) == 0 {
		return "", nil
	}

	var query string
	var args []interface{}
	var flag string
	for i, condition := range page.Conditions {
		if i == 0 {
			flag = ""
		} else {
			flag = condition.Flag
			if flag == "" {
				flag = "and"
			}
		}

		switch condition.Rule {
		case "like":
			query += fmt.Sprintf(" %s %s %s ? ", flag, condition.Field, condition.Rule)
			args = append(args, "%"+condition.Value.(string)+"%")
		case "in":
			query += fmt.Sprintf(" %s %s %s (?) ", flag, condition.Field, condition.Rule)
			args = append(args, condition.Value)
		default:
			query += fmt.Sprintf(" %s %s %s ? ", flag, condition.Field, condition.Rule)
			args = append(args, condition.Value)
		}
	}

	return query, args
}

// GetByID Find by id structure
type GetByID struct {
	ID int `json:"id" form:"id"` // 主键ID
}

type IDsReq struct {
	IDs []int `json:"ids" form:"ids"`
}

type EmptyRequest struct {
}
