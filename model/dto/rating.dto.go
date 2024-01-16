package dto

type RatingFarmerRequestDTO struct {
	Rating int    `json:"rating" validate:"required,numeric,min=1,max=5"`
	Review string `json:"review" validate:"required"`
}

type RatingFarmerResponseDTO struct {
	ID       string `json:"id"`
	FarmerID string `json:"farmer_id"`
	Rating   int    `json:"rating"`
	Review   string `json:"review"`
}

type RatingProductRequestDTO struct {
	Rating int    `json:"rating" validate:"required,numeric,min=1,max=5"`
	Review string `json:"review" validate:"required"`
}

type RatingProductResponseDTO struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Rating    int    `json:"rating"`
	Review    string `json:"review"`
}
