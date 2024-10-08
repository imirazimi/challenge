package config

import (
	"challenge/adapter/openai"
	"challenge/delivery/http"
	"challenge/repository/storage"
)

type Config struct {
	HttpServer http.Config    `koanf:"http_server"`
	Storage    storage.Config `koanf:"storage"`
	OpenAI     openai.Config  `koanf:"openai"`
}

const (
	defaultPath = "config.yml"
)

var defaultConfig = map[string]any{
	"http_server": map[string]any{
		"port": 3001,
	},
	"storage": map[string]any{
		"base":   "data",
		"laptop": "data/laptop.json",
	},
	"openai": map[string]any{
		"key":   "<put key here>",
		"model": "gpt-4o-mini",
		"role":  "user",
	},
}
