package config

import "github.com/spf13/viper"

type Config struct {
	DbHost            string `mapstructure:"DB_HOST"`
	DbUsername        string `mapstructure:"DB_USERNAME"`
	DbPort            string `mapstructure:"DB_PORT"`
	DbPassword        string `mapstructure:"DB_PASSWORD"`
	DbName            string `mapstructure:"DB_NAME"`
	JWTSecret         string `mapstructure:"JWT_SECRET"`
	Port              string `mapstructure:"PORT"`
	RajaongkirSecret  string `mapstructure:"RAJAONGKIR_KEY"`
	RajaongkirBaseUrl string `mapstructure:"RAJAONGKIR_BASEURL"`
	ShopeOriginID     string `mapstructure:"SHOP_ORIGIN_ID"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
