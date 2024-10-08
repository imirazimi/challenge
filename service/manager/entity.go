package manager

type LaptopDetail struct {
	Brand           string `json:"brand"`            // Dell
	Model           string `json:"model"`            // Inspiron
	Processor       string `json:"processor"`        // Intel Core i7-10510U
	RamCapacity     string `json:"ram_capacity"`     // 16GB
	RamType         string `json:"ram_type"`         // DDR4
	StorageCapacity string `json:"storage_capacity"` // 512GB
	BatteryStatus   string `json:"battery_status"`   // No
}

type Laptop struct {
	Key          string
	LaptopDetail LaptopDetail
}
