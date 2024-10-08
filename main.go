package main

import (
	"challenge/adapter/memcache"
	"challenge/adapter/openai"
	"challenge/config"
	"challenge/delivery/http"
	"challenge/repository/storage"
	"challenge/service"
)

func main() {
	config := Config()
	memCache := MemCache()
	openAI := OpenAI(config)
	storage := Storage(config)
	service := Service(storage, memCache, openAI)
	httpServer := HttpServer(config, service)
	httpServer.Serve()
}

func Config() *config.Config {
	return config.Load()
}
func MemCache() *memcache.MemCache {
	return memcache.New()
}
func OpenAI(config *config.Config) *openai.Adapter {
	return openai.New(config.OpenAI)
}
func Service(storage *storage.Storage, memCache *memcache.MemCache, openAI *openai.Adapter) *service.Service {
	return service.New(storage, memCache, openAI)
}
func Storage(config *config.Config) *storage.Storage {
	return storage.New(config.Storage)
}
func HttpServer(config *config.Config, service *service.Service) *http.Server {
	return http.New(config.HttpServer, service)
}
