package config

type Config struct {
	Host        string      `yaml:"host"`
	Port        string      `yaml:"port"`
	Name        string      `yaml:"name"`
	GinRelease  bool        `yaml:"release"`
	{{if .mysql}}
	MysqlConfig MysqlConfig `yaml:"mysqlConfig"`
	{{end}}

	{{if .redis}}
	RedisConfig RedisConfig `yaml:"redisConfig"`
	{{end}}
}

{{if .mysql}}
type MysqlConfig struct {
	Ip              string `yaml:"ip"`
	Port            int    `yaml:"port"`
	Pwd             string `yaml:"pwd"`
	User            string `yaml:"user"`
	ConnectPoolSize int    `yaml:"connectPoolSize"`
	SetLog          bool   `yaml:"setLog"`
}

{{end}}

{{if .redis}}
type RedisConfig struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
	Pwd  string `yaml:"pwd"`
	Db   []int  `yaml:"db"`
}
{{end}}