package entity

import "time"

type User struct {
	ID       int64  `json:"id" gorm:"primary_key;column:id"` // 主键
	UserName string `json:"userName" gorm:"column:username"` // 用户名
	Password string `json:"password" gorm:"column:password"` // 密码
	Email    string `json:"email" gorm:"column:email"`       // 邮箱
}

type Article struct {
	ID         int64     `json:"id" gorm:"primary_key;column:id"`                     // 主键
	Title      string    `json:"title" gorm:"column:title"`                           // 文章标题
	UserID     int64     `json:"userId" gorm:"column:user_id"`                        // 用户ID
	Content    string    `json:"content" gorm:"column:content"`                       // 文章内容
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;autoCreateTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;autoUpdateTime"` // 修改时间
}

type Comment struct {
	ID         int64     `json:"id" gorm:"primary_key;column:id"`                     // 评论ID
	UserId     int64     `json:"userId" gorm:"column:user_id"`                        // 评论用户ID
	ArticleID  int64     `json:"articleId" gorm:"column:article_id"`                  // 评论文章ID
	Content    string    `json:"content" gorm:"column:content"`                       // 评论内容
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;autoCreateTime"` // 创建时间
}
