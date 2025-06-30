package models

type Configs struct {
	AuthParams     AuthParams     `json:"auth_params"`
	LogParams      LogParams      `json:"log_params"`
	AppParams      AppParams      `json:"app_params"`
	PostgresParams PostgresParams `json:"postgres_params"`
}
type AuthParams struct {
	JwtSecretKey  string `json:"jwt_secret_key"`
	JwtTtlMinutes int    `json:"jwt_ttl_minutes"`
}


type LogParams struct {
	LogDirectory      string `json:"LogDirectory"`
	LogInfo           string `json:"LogInfo"`
	LogError          string `json:"LogError"`
	LogWarn           string `json:"LogWarn"`
	LogDebug          string `json:"LogDebug"`
	MaxSizeMegabytes  int    `json:"MaxSizeMegabytes"`
	MaxBackups        int    `json:"MaxBackups"`
	MaxAgeDays        int    `json:"MaxAgeDays"`
	Compress          bool   `json:"Compress"`
	LocalTime         bool   `json:"LocalTime"`
}
type AppParams struct {
	ServerURL  string `json:"server_url"`
	ServerName string `json:"server_name"`
	AppVersion string `json:"app_version"`
	PortRun    string `json:"port_run"`
	GinMode    string `json:"gin_mode"`
}

type PostgresParams struct {
	User     string `json:"user"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}
