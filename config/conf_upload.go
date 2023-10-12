package config

type Upload struct {
	Size      int    `yaml:"size" json:"size" default:"5"`
	ImagePath string `yaml:"image_path" json:"image_path" default:"uploads/images"`
	FilePath  string `yaml:"file_path" json:"file_path" default:"uploads/files"`
	AudioPath string `yaml:"audio_path" json:"audio_path" default:"uploads/audios"`
}
