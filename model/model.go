package model

// import (
//     "time"
//     "gorm.io/gorm"
// )

type TypechoContent struct {
	Cid          uint   `gorm:"column:cid;primaryKey;autoIncrement"`              // 内容ID（自增主键）
	Title        string `gorm:"column:title;type:varchar(200);not null"`          // 标题
	Slug         string `gorm:"column:slug;type:varchar(200);index"`              // 友好URL slug
	Text         string `gorm:"column:text;type:longtext;not null"`               // 正文内容（支持HTML/Markdown）
	Order        int    `gorm:"column:order;type:int;default:0"`                  // 排序权重
	AuthorId     uint   `gorm:"column:authorId;type:int unsigned;not null"`       // 作者ID
	Template     string `gorm:"column:template;type:varchar(32)"`                 // 自定义模板
	Type         string `gorm:"column:type;type:varchar(16);default:'post'"`      // 内容类型（post/page/draft）
	Status       string `gorm:"column:status;type:varchar(16);default:'publish'"` // 状态（publish/private）
	Password     string `gorm:"column:password;type:varchar(32)"`                 // 访问密码（加密后）
	CommentsNum  int    `gorm:"column:commentsNum;type:int;default:0"`            // 评论数（冗余计数）
	AllowComment int    `gorm:"column:allowComment;type:tinyint(1);default:1"`    // 允许评论（0/1）
	AllowPing    int    `gorm:"column:allowPing;type:tinyint(1);default:1"`       // 允许Ping（0/1）
	AllowFeed    int    `gorm:"column:allowFeed;type:tinyint(1);default:1"`       // 允许订阅（0/1）
	Parent       uint   `gorm:"column:parent;type:int unsigned;default:0"`        // 父内容ID（用于页面层级）
	Created      int64  `gorm:"column:created;type:int;index"`                    // 创建时间（Unix时间戳）
	Modified     int64  `gorm:"column:modified;type:int;index"`                   // 修改时间（Unix时间戳）
}

func (TypechoContent) TableName() string {
	return "typecho_contents"
}
