package setting

type Config struct {
	Server ServerConfig  `mapstructure:"server"`
	Mysql  MysqlConfig   `mapstructure:"mysql"`
	Redis  RedisConfig   `mapstructure:"redis"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"` // "debug", "release", "test"
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	MaxIdle  int    `mapstructure:"maxIdleConns"`
	MaxOpen  int    `mapstructure:"maxOpenConns"`
	MaxLife  int    `mapstructure:"maxConnLifetime"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}
