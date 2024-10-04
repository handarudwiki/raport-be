package config

type Config struct {
	Database Database
	App      App
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type App struct {
	Host    string
	Port    string
	Name    string
	Version string
	Env     string
}
