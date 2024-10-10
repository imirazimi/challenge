package manager

import (
	"challenge/contract"
	"context"
	"encoding/json"
	"log"
)

//go:generate mockery --name Repo
type Repo interface {
	AddLaptop(ctx context.Context, laptop []Laptop) error
	GetAllLaptop(ctx context.Context) ([]Laptop, error)
}

type Service struct {
	repo  Repo
	cache contract.Cache
	agent contract.Agent
}

// Create Service And Seed Cache With Persisted Data
func New(repo Repo, cache contract.Cache, agent contract.Agent) Service {
	service := Service{repo: repo, cache: cache, agent: agent}
	ctx := context.Background()
	laptops, err := service.repo.GetAllLaptop(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, laptop := range laptops {
		err = service.cache.Set(ctx, laptop.Key, laptop.LaptopDetail)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return Service{repo: repo, cache: cache, agent: agent}
}

// Integrate New LaptopDetail With LLM Agent Model
// If All LaptopDetail(s) Return From Cache If Exist
func (s Service) IntegrateLaptop(ctx context.Context, req IntegrateLaptopRequest) (IntegrateLaptopResponse, error) {

	newsKeys := []string{}
	olds := []LaptopDetail{}
	for _, data := range req.Data {
		d, err := s.cache.Get(ctx, data)
		if err != nil {
			newsKeys = append(newsKeys, data)
		} else {
			laptopDetail := d.(LaptopDetail)
			olds = append(olds, laptopDetail)
		}
	}
	if len(newsKeys) == 0 {
		return IntegrateLaptopResponse{Data: olds}, nil
	}

	news := []LaptopDetail{}
	newsJSON, err := s.agent.JSON(ctx, newsKeys, LaptopDetailFormat)
	if err != nil {
		return IntegrateLaptopResponse{}, err
	}

	err = json.Unmarshal([]byte(newsJSON), &news)
	if err != nil {
		return IntegrateLaptopResponse{}, err
	}

	newLaptops := []Laptop{}
	for i, _ := range news {
		laptop := Laptop{
			Key:          newsKeys[i],
			LaptopDetail: news[i],
		}
		newLaptops = append(newLaptops, laptop)
	}
	err = s.repo.AddLaptop(ctx, newLaptops)
	if err != nil {
		return IntegrateLaptopResponse{}, err
	}
	for i, _ := range news {
		err = s.cache.Set(ctx, newsKeys[i], news[i])
		if err != nil {
			return IntegrateLaptopResponse{}, err
		}
	}
	all := []LaptopDetail{}
	all = append(all, olds...)
	all = append(all, news...)
	return IntegrateLaptopResponse{Data: all}, nil
}

// Return All LaptopDetail(s) From Cache
// Cache Is Always Sync With Persist Data Layer
func (s Service) GetAllLaptop(ctx context.Context, _ GetAllLaptopRequest) (GetAllLaptopResponse, error) {
	data, err := s.cache.GetAll(ctx)
	if err != nil {
		return GetAllLaptopResponse{}, err
	}
	all := []LaptopDetail{}
	for _, d := range data {
		laptopDetail := d.(LaptopDetail)
		all = append(all, laptopDetail)
	}
	return GetAllLaptopResponse{Data: all}, nil
}
