package service

import (
	"challenge/adapter/memcache"
	"challenge/adapter/openai"
	"challenge/repository/storage"
	"challenge/service/manager"
)

type Service struct {
	ManagerSvc manager.Service
}

func New(storage *storage.Storage, memCache *memcache.MemCache, openAI *openai.Adapter) *Service {
	managerSvc := manager.New(storage, memCache, openAI)
	return &Service{
		ManagerSvc: managerSvc,
	}
}
