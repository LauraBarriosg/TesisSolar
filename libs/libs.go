package libs

import (
	"log"
	"os"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
var DB *gorm.DB
type DbConfig struct {
	Host      string		`mapstructure:"DB_HOST"`
	Port         string		`mapstructure:"DB_PORT"`
	Database     string		`mapstructure:"DB_DATABASE"`
	User         string		`mapstructure:"DB_USER"`
	Password     string		`mapstructure:"DB_PASSWORD"`
	Charset      string		`mapstructure:"DB_CHARSET"`
	MaxIdleConns int		`mapstructure:"DB_MAX_ID"`
	MaxOpenConns int		`mapstructure:"DB_MAX_OPEN"`
	print_log    bool		`mapstructure:"DB_LOG"`
}


// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, ConfigName string) (DbConfig DbConfig, err error) {
	
	config.AddConfigPath("./app/")   // path to look for the config file in
	config.AddConfigPath("/go/src/app")  // call multiple times to add many search paths
	config.AddConfigPath(".")               // optionally look for config in the working directory
	config.SetConfigName(ConfigName)
	config.SetConfigType("env")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if config.GetString("DB_HOST") == "" {
		fmt.Println("falta host en el archivo " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("DB_DATABASE") == "" {
		fmt.Println("falta database en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("DB_USER") == "" {
		fmt.Println("falta user en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("DB_PASSWORD") == "" {
		fmt.Println("falta password en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	err = config.Unmarshal(&DbConfig)
	return
}






























/*import (
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//config "github.com/spf13/viper"
)

var DB *gorm.DB

type DbConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	print_log    bool
}


func Configure() DbConfig {
	
	response := DbConfig{
		Host:"postgres",
		Port:"15432",
		Database:"Tesis_Estacion",
		User:"EstacionSolar",
		Password:"Tesis",
		Charset:"utf8",
		MaxIdleConns:10,
		MaxOpenConns:100,
		print_log:true,
		
	}


	return response
}

func Configure(ConfigName string, Extension string) DbConfig {
	config.SetConfigName(ConfigName)
	config.SetConfigType(Extension)
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if config.GetString("default.host") == "" {
		fmt.Println("falta host en el archivo " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.database") == "" {
		fmt.Println("falta database en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.user") == "" {
		fmt.Println("falta user en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.password") == "" {
		fmt.Println("falta password en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	
	response := DbConfig{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
		config.GetString("default.charset"),
		config.GetInt("default.MaxIdleConns"),
		config.GetInt("default.MaxOpenConns"),
		config.GetBool("default.sql_log"),
	}


	return response
}*/