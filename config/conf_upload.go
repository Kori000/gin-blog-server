package config

type Upload struct {
	Size      int    `yaml:"size" json:"size"`
	ImagePath string `yaml:"image_path" json:"image_path"`
	FilePath  string `yaml:"file_path" json:"file_path"`
	AudioPath string `yaml:"audio_path" json:"audio_path"`
}
