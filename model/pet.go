package model

type Pet struct {
	PetID    int    `json:"petid" form:"petid"`
	Category string `json:"category" form:"category"`
	Name     string `json:"name" form:"name"`
	Umur     int    `json:"umur" form:"umur"`
	Berat    int    `json:"berat" form:"berat"`
}
