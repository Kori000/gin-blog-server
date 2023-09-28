package models

type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  // 子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` // 父级评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                               // 父级评论
	Content            string          `gorm:"size:256" json:"content"`                         // 文章内容
	DiggCount          int             `json:"digg_count"`                                      // 点赞量
	CommentCount       int             `json:"comment_count"`                                   // 评论量
	Article            ArticleModel    `gorm:"foreignKey:ArticleID" json:"-"`                   // 关联的文章
	ArticleID          uint            `json:"article_id"`                                      // 文章id
	User               UserModel       `json:"user"`                                            // 评论的用户
	UserID             uint            `json:"user_id"`                                         // 评论的用户id
}
