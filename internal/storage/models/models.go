package models

import "database/sql"

// File 结构体
type File struct {
	ID        int
	Name      string
	Path      string // 文件实际路径
	Tags      []Tag  `gorm:"many2many:file_tags;"`  // GORM 标签关联
	Boxes     []Box  `gorm:"many2many:file_boxes;"` // 多对多关联 Boxes
	Important bool
}

// Box 结构体
type Box struct {
	ID       int
	Name     string
	ParentID *int   // 允许为空，表示根层级
	Children []Box  `gorm:"foreignkey:ParentID"`   // 子分类
	Files    []File `gorm:"many2many:file_boxes;"` // 关联文件
}

type Tag struct {
	ID    int
	Name  string
	Color sql.NullString
}
