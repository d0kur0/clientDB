package configMgr

type ServerConfig struct {
	Port  int         `json:"port"`
	Mysql MysqlConfig `json:"mysql"`
	Redis RedisConfig `json:"redis"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Dbname   int    `json:"dbname"`
}
