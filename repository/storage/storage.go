package storage

import (
	"challenge/service/manager"
	"context"
	"encoding/json"
	"os"
)

type Config struct {
	Base   string `koanf:"base"`
	Laptop string `koanf:"laptop"`
}

type Storage struct {
	config Config
}

func New(config Config) *Storage {
	_ = os.Mkdir(config.Base, 0600)
	return &Storage{config: config}
}

func (s *Storage) AddLaptop(ctx context.Context, laptop []manager.Laptop) error {
	laptops, err := s.GetAllLaptop(ctx)
	if err != nil {
		return err
	}
	laptops = append(laptops, laptop...)
	file, err := os.OpenFile(s.config.Laptop, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
	}()
	buffer, err := json.Marshal(&laptops)
	if err != nil {
		return err
	}
	_, err = file.Write(buffer)
	if err != nil {
		return err
	}
	return nil
}
func (s *Storage) GetAllLaptop(_ context.Context) ([]manager.Laptop, error) {
	laptops := []manager.Laptop{}
	file, err := os.Open(s.config.Laptop)
	if err != nil {
		return []manager.Laptop{}, nil
	}
	defer func() {
		err = file.Close()
	}()
	err = json.NewDecoder(file).Decode(&laptops)
	if err != nil {
		return []manager.Laptop{}, err
	}
	return laptops, nil
}
