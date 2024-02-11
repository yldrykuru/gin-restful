// configs/config.go
package configs

import (
	"github.com/spf13/viper"
)

// Config struct'ı, projenizin genel yapılandırma ayarlarını içerir
type Config struct {
	Port        string
	DatabaseURL string
	// Diğer yapılandırma ayarları buraya eklenir
}

var config *Config

// init fonksiyonu, projenin başlangıcında yapılandırma ayarlarını yükler
func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// Yapılandırma ayarlarını yükle
	config = &Config{
		Port:        viper.GetString("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
		// Diğer yapılandırma ayarları buraya eklenir
	}
}

// GetConfig fonksiyonu, yapılandırma ayarlarını döndürür
func GetConfig() *Config {
	return config
}
