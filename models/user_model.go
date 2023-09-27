package models

import "gin-blog-server/models/ctype"

type UserModel struct {
	MODEL
	NickName       string           `gorm:"size:36" json:"nick_name"`                                                         // 昵称
	Username       string           `gorm:"size:36" json:"user_anme"`                                                         // 用户名(登录账号)
	Password       string           `gorm:"size:128" json:"password"`                                                         // 密码
	Avatar         uint             `gorm:"size:265" json:"avatar"`                                                           // 头像
	Email          string           `gorm:"size:128" json:"email"`                                                            // 邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                                                               // 电话
	Addr           string           `gorm:"size:64" json:"addr"`                                                              // 地址
	Token          string           `gorm:"size:64" json:"token"`                                                             // 其他平台的唯一id
	IP             string           `gorm:"size:28" json:"ip"`                                                                // ip地址
	Role           ctype.Role       `gorm:"size:4,default:1" json:"role"`                                                     // 角色
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`                                              // 注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:AuthID" json:"-"`                                                       // 发布的文章列表(一对多)
	CollectsModels []ArticleModel   `gorm:"many2many:auth2_collects;joinForeignKey:AuthID;joinReferences:ArticleID" json:"-"` // 收藏的文章列表(多对多)
}
