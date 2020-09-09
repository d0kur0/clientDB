package configMgr

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Password string `json:"password"`
	Dbname string `json:"dbname"`
}

type ServerConfig struct {
	Port int `json:"port"`
	Mysql MysqlConfig `json:"mysql"`
}