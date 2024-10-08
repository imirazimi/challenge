package manager

type IntegrateLaptopRequest struct {
	Data []string `json:"data"`
}

type IntegrateLaptopResponse struct {
	Data []LaptopDetail `json:"data"`
}

type GetAllLaptopRequest struct {
}

type GetAllLaptopResponse struct {
	Data []LaptopDetail `json:"data"`
}
