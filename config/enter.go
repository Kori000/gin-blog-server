package config

// 程序启动时, 会将所有 Config 保存在 Global
type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	Email    Email    `yaml:"email"`
	QiNiu    QiNiu    `yaml:"qiniu"`
	JWT      JWT      `yaml:"jwt"`
}
