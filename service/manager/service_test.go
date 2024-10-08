package manager

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrateLaptop(t *testing.T) {
	mockRepo := NewMockRepo(t)
	mockAgent := NewMockAgent(t)
	mockCache := NewMockCache(t)

	t.Run("Cache Hit", func(t *testing.T) {
		ctx := context.Background()
		mockRepo.EXPECT().GetAllLaptop(ctx).Return([]Laptop{}, nil).Once()
		svc := New(mockRepo, mockCache, mockAgent)
		req := IntegrateLaptopRequest{
			Data: []string{"Laptop: Dell Inspiron; Processor i7-10510U ; RAM 16GB; 512GB SSD Missing battery"},
		}

		mockCache.EXPECT().Get(ctx, req.Data[0]).Return(LaptopDetail{Brand: "Dell", Model: "Inspiron", Processor: "i7-10510U", RamCapacity: "16GB", RamType: "", StorageCapacity: "512GB", BatteryStatus: "Missing battery"}, nil).Once()

		resp, err := svc.IntegrateLaptop(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, IntegrateLaptopResponse{
			Data: []LaptopDetail{
				{Brand: "Dell", Model: "Inspiron", Processor: "i7-10510U", RamCapacity: "16GB", RamType: "", StorageCapacity: "512GB", BatteryStatus: "Missing battery"},
			},
		}, resp)
	})
}

func TestGetAllLaptop(t *testing.T) {
	mockRepo := NewMockRepo(t)
	mockAgent := NewMockAgent(t)
	mockCache := NewMockCache(t)

	t.Run("Cache Hit", func(t *testing.T) {
		ctx := context.Background()
		req := GetAllLaptopRequest{}
		mockRepo.EXPECT().GetAllLaptop(ctx).Return([]Laptop{
			{Key: "Dell Inspiron i7-10510U 16GB 512GB Missing battery", LaptopDetail: LaptopDetail{Brand: "Dell", Model: "Inspiron", Processor: "i7-10510U", RamCapacity: "16GB", RamType: "", StorageCapacity: "512GB", BatteryStatus: "Missing battery"}},
			{Key: "Lenovo ThinkPad i7-7700H 8GB DDR3 512GB Battery Replaced", LaptopDetail: LaptopDetail{Brand: "Lenovo", Model: "ThinkPad", Processor: "i7-7700H", RamCapacity: "8GB", RamType: "DDR3", StorageCapacity: "512GB", BatteryStatus: "Battery Replaced"}},
			{Key: "HP EliteBook i7-3520M 16GB DDR3 512GB Battery Healthy", LaptopDetail: LaptopDetail{Brand: "HP", Model: "EliteBook", Processor: "i7-3520M", RamCapacity: "16GB", RamType: "DDR3", StorageCapacity: "512GB", BatteryStatus: "Battery Healthy"}},
		}, nil).Once()
		mockCache.EXPECT().Set(ctx, "Dell Inspiron i7-10510U 16GB 512GB Missing battery", LaptopDetail{Brand: "Dell", Model: "Inspiron", Processor: "i7-10510U", RamCapacity: "16GB", RamType: "", StorageCapacity: "512GB", BatteryStatus: "Missing battery"}).Return(nil).Once()
		mockCache.EXPECT().Set(ctx, "Lenovo ThinkPad i7-7700H 8GB DDR3 512GB Battery Replaced", LaptopDetail{Brand: "Lenovo", Model: "ThinkPad", Processor: "i7-7700H", RamCapacity: "8GB", RamType: "DDR3", StorageCapacity: "512GB", BatteryStatus: "Battery Replaced"}).Return(nil).Once()
		mockCache.EXPECT().Set(ctx, "HP EliteBook i7-3520M 16GB DDR3 512GB Battery Healthy", LaptopDetail{Brand: "HP", Model: "EliteBook", Processor: "i7-3520M", RamCapacity: "16GB", RamType: "DDR3", StorageCapacity: "512GB", BatteryStatus: "Battery Healthy"}).Return(nil).Once()
		svc := New(mockRepo, mockCache, mockAgent)
		mockCache.EXPECT().GetAll(ctx).Return([]any{
			LaptopDetail{Brand: "Dell", Model: "Inspiron", Processor: "i7-10510U", RamCapacity: "16GB", RamType: "", StorageCapacity: "512GB", BatteryStatus: "Missing battery"},
			LaptopDetail{Brand: "Lenovo", Model: "ThinkPad", Processor: "i7-7700H", RamCapacity: "8GB", RamType: "DDR3", StorageCapacity: "512GB", BatteryStatus: "Battery Replaced"},
			LaptopDetail{Brand: "HP", Model: "EliteBook", Processor: "i7-3520M", RamCapacity: "16GB", RamType: "DDR3", StorageCapacity: "512GB", BatteryStatus: "Battery Healthy"},
		}, nil).Once()

		resp, err := svc.GetAllLaptop(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, GetAllLaptopResponse{
			Data: []LaptopDetail{
				{Brand: "Dell", Model: "Inspiron", Processor: "i7-10510U", RamCapacity: "16GB", RamType: "", StorageCapacity: "512GB", BatteryStatus: "Missing battery"},
				{Brand: "Lenovo", Model: "ThinkPad", Processor: "i7-7700H", RamCapacity: "8GB", RamType: "DDR3", StorageCapacity: "512GB", BatteryStatus: "Battery Replaced"},
				{Brand: "HP", Model: "EliteBook", Processor: "i7-3520M", RamCapacity: "16GB", RamType: "DDR3", StorageCapacity: "512GB", BatteryStatus: "Battery Healthy"},
			},
		}, resp)
	})
}
