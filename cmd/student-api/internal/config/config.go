package config

type HttpServer struct {
	Address string
}

type Config struct {
	Env          string `yaml: "env"`
	Storage_path string `yaml: "storage_path" env-required:true`
	HttpServer   `yaml:"http_server"`
}
