package config

type DBConfig struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Url      string `yaml:"url"`
}

type Config struct {
	dbcfg DBConfig `yaml:"db"`
}
