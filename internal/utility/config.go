package utility

import "github.com/spf13/viper"

type Config struct {
	DBDriver string  `mapstructure:"DB_DRIVER"`
	DSN      string  `mapstructure:"POSTGRES_DSN"`
	Address  string  `mapstructure:"ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("app")
    viper.SetConfigType("env")
    
    viper.AutomaticEnv()
    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}
