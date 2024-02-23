package config

type Config struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Name       string `yaml:"name"`
	GinRelease bool   `yaml:"release"`
}
