package tomlConfig

type SL struct {
	System system
	Elk elk
}

type system struct {
	HttpServer string	`toml:"httpServer"`
	EndPoint string	`toml:"endPoint"`
	Scripts string	`toml:"scripts"`
	SlowLog string	`toml:"slowLog"`
	LogFile string	`toml:"logFile"`
}

type elk struct {
	EndPoint []string	`toml:"endPoint"`
	Index string	`toml:"index"`
	Product string	`toml:"product"`
	Cluster string	`toml:"cluster"`
	Role string		`toml:"role"`
	HostName string	`toml:"hostName"`
}














