package models

import "gin-blog-server/models/ctype"

type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"`                   // 文章表题
	Abstract      string         `json:"abstract"`                               //	文章介绍
	Content       string         `json:"content"`                                // 文章内容
	LookCount     int            `json:"lookcount"`                              // 浏览量
	CommentCount  int            `json:"comment_count"`                          // 评论量
	DiggCount     int            `json:"digg_count"`                             //	点赞量
	CollectsCount int            `json:"collects_count"`                         //	收藏量
	TagModels     []TagModel     `gorm:"many2many:article_tag" json:"tagmodels"` // 文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`          //	文章评论列表
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`             // 文章作者
	UserID        uint           `json:"user_id"`                                //	用户id
	Category      string         `gorm:"size:20" json:"category"`                // 文章分类
	Source        string         `json:"source"`                                 //	文章来源
	Link          string         `json:"link"`                                   //	原文链接
	Cover         BannerModel    `json:"-"`                                      // 文章封面
	CoverID       uint           `json:"cover_id"`                               // 文章封面id
	NickName      string         `gorm:"size:42" json:"nickname"`                //	发布文章的用户昵称
	CoverPath     string         `json:"cover_path"`                             //	文章的封面
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`        //	文章标签
}
