package model

type Pet struct {
	PetID    int    `json:"petid" form:"petid"`
	Kategori string `json:"kategori" form:"kategori"`
	Name     string `json:"name" form:"name"`
	Umur     int    `json:"umur" form:"umur"`
	Berat    int    `json:"berat" form:"berat"`
}
