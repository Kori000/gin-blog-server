package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at" json:"created_at" default:"2023-02-12"`
	Filing      string `yaml:"filing" json:"filing" default:"湘ICP备xxxxx号"`
	Title       string `yaml:"title" json:"title" default:"个人博客"`
	QQImage     string `yaml:"qq_image" json:"qq_image" default:"https://avatars.githubusercontent.com/u/9919?s=80&v=4"`
	Version     string `yaml:"version" json:"version" default:"0.0.1"`
	Email       string `yaml:"email" json:"email" default:"kexin@korix.com"`
	WechatImage string `yaml:"wechat_image" json:"wechat_image" default:"https://avatars.githubusercontent.com/u/9919?s=80&v=4"`
	Name        string `yaml:"name" json:"name" default:"KORIX"`
	Job         string `yaml:"job" json:"job" default:"FE"`
	Addr        string `yaml:"addr" json:"addr" default:"Earth"`
	Slogan      string `yaml:"slogan" json:"slogan" default:"~"`
	SloganEn    string `yaml:"slogan_en" json:"slogan_en" default:"~~"`
	Web         string `yaml:"web" json:"web" default:"https://github.com/Kori000"`
	BiliBiliUrl string `yaml:"bilibili_url" json:"bilibili_url" default:"https://github.com/Kori000"`
	GiteeUrl    string `yaml:"gitee_url" json:"gitee_url" default:"https://github.com/Kori000"`
	GithubUrt   string `yaml:"github_urt" json:"github_urt" default:"https://github.com/Kori000"`
}
