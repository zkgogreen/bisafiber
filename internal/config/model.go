package config

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Name string `json:"database"`
	User string `json:"username"`
	Pass string `json:"password"`
	Tz   string `json:"timezone"`
}
