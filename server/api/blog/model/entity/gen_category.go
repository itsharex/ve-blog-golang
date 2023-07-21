package entity

import "time"

// TableNameCategory return the table name of <category>
const TableNameCategory = "category"

// Category mapped from table <category>
type Category struct {
	ID           int       `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	CategoryName string    `gorm:"column:category_name;type:varchar(20);not null;comment:分类名" json:"category_name"` // 分类名
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`         // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                  // 更新时间
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
