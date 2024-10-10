package manager

const (
	LaptopDetailFormat = `
[
	{
		"brand":"<if brand name doesn't exist then explore brand based on model name>",
		"model":"<just return model name without brand name>",
		"processor":"<explore and find processor brand if not exists>",
		"ram_capacity":"<just extract capacity without decimal part and with unit>",
		"ram_type":"<if data is empty, then explore and find ram type based on processor model and return Just 'DDR4' or 'DDR3'>",
		"storage_capacity":"<just extract capacity without decimal part and with unit>",
		"battery_status":"<if data is empty return Yes.if battery is corrupted, damaged or unused return No otherwise Yes>"
	}
]`
)

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
