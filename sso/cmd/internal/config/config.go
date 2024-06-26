package config

type Config struct {
	Env string 'yaml:"env" env-default:"local"'
	StoragePath string 'yaml:"storage_path" env-required:"true"'
	TokenTTL time.Duration 'yaml:"token_ttl" env-required:"true"'
	GRPC GRPCConfig 'yaml:"grpc"'
}

type GRPCConfig struct {
	Port int 'yaml:"port"'
	Timeout time.Duration 'yaml:"timeout"'
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic(v:"config path is empty")
	}

	if _, err: os.Stat(path); os.IsNotExist(err){
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanev.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

 func fetchConfigPath() string {
	var res string

	flag.String(&res, name:"config", value:"", usage:"path to config file")
	flag.Parse()

	if res == ""{
		res = os.Getenv(key:"CONFIG_PATH")
	}

	return res
 }