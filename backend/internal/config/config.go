package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Description string `mapstructure:"description"`
}

type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type ApiConfig struct {
	RateLimit string `mapstructure:"rate_limit"`
	Timeout   string `mapstructure:"timeout"`
}

type Config struct {
	// Database (แยก environment variables มาใช้สำหรับ DB)
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`

	// Redis
	RedisURL string `mapstructure:"REDIS_URL"`

	// JWT
	JwtSecret string `mapstructure:"JWT_SECRET"`

	// App (ใช้จาก YAML)
	App AppConfig `mapstructure:"app"`

	// Logging (ใช้จาก YAML)
	Logging LoggingConfig `mapstructure:"logging"`

	// API (ใช้จาก YAML)
	Api ApiConfig `mapstructure:"api"`
}

func LoadConfig() (*Config, error) {
	// 1️⃣ โหลดค่าจาก .env
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: failed to load .env file: %v\n", err)
	} else {
		fmt.Println("[INFO] .env file loaded")
	}

	// 2️⃣ ตั้งค่าให้ Viper โหลดค่าจากไฟล์ YAML
	viper.SetConfigName("app")        // กำหนดชื่อไฟล์ config (app.yaml)
	viper.AddConfigPath("../configs") // ระบุ path ที่เก็บไฟล์
	viper.SetConfigType("yaml")       // ระบุประเภทของไฟล์เป็น YAML

	// 3️⃣ ลองอ่าน config จาก YAML
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Warning: no config file found: %v\n", err)
	}

	// 4️⃣ ทำให้ Viper อ่านค่า environment variables
	viper.AutomaticEnv()

	// 5️⃣ กำหนดให้ค่าจาก Environment Variables Override ค่าจาก YAML สำหรับเฉพาะค่าของ DB
	// ค่าที่จะถูก override เป็นค่าจาก environment variables (เฉพาะ DB, REDIS, JWT)
	_ = viper.BindEnv("DB_HOST")
	_ = viper.BindEnv("DB_PORT")
	_ = viper.BindEnv("DB_USER")
	_ = viper.BindEnv("DB_PASS")
	_ = viper.BindEnv("DB_NAME")
	_ = viper.BindEnv("REDIS_URL")
	_ = viper.BindEnv("JWT_SECRET")

	// 6️⃣ แสดงค่าทั้งหมดที่โหลดมา
	fmt.Println("[DEBUG] All settings:", viper.AllSettings())

	// 7️⃣ Unmarshal ค่าที่โหลดมาเป็น struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	// แสดงค่าที่ unmarshaled
	fmt.Printf("[INFO] Final config: %+v\n", config)
	return &config, nil
}
