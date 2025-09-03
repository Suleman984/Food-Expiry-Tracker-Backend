package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseUrl       string
	SupabaseAnonKey   string
	SupabaseServiceKey string
	SMTPEmail         string
	SMTPPassword      string
}

func LoadConfig() Config {
	_ = godotenv.Load() // load .env file if exists

	cfg := Config{
		SupabaseUrl:        os.Getenv("SUPABASE_URL"),
		SupabaseAnonKey:    os.Getenv("SUPABASE_ANON_KEY"),
		SupabaseServiceKey: os.Getenv("SUPABASE_SERVICE_KEY"),
		SMTPEmail:          os.Getenv("SMTP_EMAIL"),
		SMTPPassword:       os.Getenv("SMTP_PASSWORD"),
	}

	if cfg.SupabaseUrl == "" || cfg.SupabaseServiceKey == "" {
		log.Fatal("Missing Supabase credentials")
	}

	return cfg
}
