package config

func New() *Config {
	return &Config{
		Secrets: new(Secrets),
	}
}

type Config struct {
	RabbitMq *RabbitMq `json:"rabbitMq"`
	Secrets  *Secrets  `json:"secrets"`
}

type RabbitMq struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Secrets struct {
	RabbitMqCredentials *RabbitMqCredentials `json:"rabbitMqCredentials"`
}

type RabbitMqCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
