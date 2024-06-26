module github.com/IvkinNikita/GoAuthorization

go 1.21

require (
	github.com/IvkinNikita/GoAuthorization/sso/internal/config v0.0.0-20240626080539-94d65a5853b0
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace github.com/IvkinNikita/GoAuthorization/sso/internal/config => ./sso/internal/config

